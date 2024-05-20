/* dbBuilder.go
2024, cdfisher
----------------
Loads OSRS cache dumps either retrieved from github.com/abextm/osrs-cache
or dumped using the dumper from github.com/abextm/osrs-flatcache into a SQLite3 DB.

Currently, this supports loading from the item_defs, npc_defs, and object_defs directories.

Table creations and insertions are based on keys found in Abex's dump of cache 221.7
(2024-05-15-rev221)
*/

/*
// TODO
dbBuilder TODOs:
----------------
- Take cache version/path as CL args so this can just be provided as an executable
- Add handling for when new keys are included in definition entries: Add a table/tables for new keys
- Add support for other objects in cache: dbtables, param_defs, ???
- Maybe reorder columns to match appearance in defs/group them a bit more sensibly

*/

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"log"
	"os"
	"path/filepath"
	"time"
)

func SliceTextStr(slice []string) string {
	str, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("Error marshalling slice %s\n", slice)
	}
	return fmt.Sprintf("%v", string(str))
}

func SliceTextInt(slice []int) string {
	str, err := json.Marshal(slice)
	if err != nil {
		fmt.Printf("Error marshalling slice %v\n", slice)
	}
	return fmt.Sprintf("%v", string(str))
}

func MapToStr(mapInput map[string]interface{}) string {
	mapStr, err := json.Marshal(mapInput)
	if err != nil {
		fmt.Printf("Error marshalling map %s\n", mapInput)
	}
	return fmt.Sprintf("%v", string(mapStr))
}

func getFileNames(dir string) []string {
	names := make([]string, 0, 100)
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fileName := filepath.Base(path)
		if !info.IsDir() && info.Size() > 0 {
			names = append(names, fileName)
		}
		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
	return names
}

func initializeDB(dbfile string) *sql.DB {
	// load schema.sql into string
	creationStatement, err := os.ReadFile("schema.sql")
	if err != nil {
		log.Fatal("Error reading schema file: ", err)
	}
	creationString := string(creationStatement)

	// Open/create db file
	database, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		log.Fatal("Could not open or create db file: ", err)
	}

	// Create table based on schema.sql
	_, err = database.Exec(creationString)
	if err != nil {
		log.Fatal("Could not execute creation statement: ", err)
	}

	return database
}

func insertItemData(cachePath string, database *sql.DB) {
	// get all files in /item_defs
	itemFiles := getFileNames(fmt.Sprintf("%s\\item_defs", cachePath))

	// for each file loop through
	for i := range itemFiles {
		// if file is not empty: (checked with getFileNames)
		// read file
		fileBytes, err := os.ReadFile(fmt.Sprintf("%s\\item_defs\\%s", cachePath, itemFiles[i]))
		if err != nil {
			fmt.Printf("Error reading item file %s to bytes: %s\n", itemFiles[i], err)
		}
		// unmarshal into an ItemEntry "def"
		def := ItemEntry{}
		if err = json.Unmarshal(fileBytes, &def); err != nil {
			fmt.Printf("Error unmarshalling item file %s : %s\n", itemFiles[i], err)
		}

		// SQL time
		statement := "INSERT OR REPLACE INTO items (id, name, examine, resize_x, resize_y, resize_z, xan2d, yan2d, zan2d, cost, is_tradable, stackable, inventory_model, wear_pos_1, wear_pos_2, wear_pos_3, members, zoom_2d, x_offset_2d, y_offset_2d, ambient, contrast, options, interface_options, male_model_0, male_model_1, male_model_2, male_offset, male_head_model, male_head_model_2, female_model_0, female_model_1, female_model_2, female_offset, female_head_model, female_head_model_2, noted_id, noted_template, team, weight, shift_click_drop_index, bought_id, bought_template_id, placeholder_id, placeholder_template_id, color_find, color_replace, params, count_co, count_obj, texture_find, texture_replace, category) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		preparedStatement, err := database.Prepare(statement)
		if err != nil {
			fmt.Printf("Error preparing item insertion statement for item %s : %s\n", itemFiles[i], err)
			log.Fatal(err)
		}

		// I don't love this
		_, err = preparedStatement.Exec(def.ID, def.Name, def.Examine, def.ResizeX, def.ResizeY, def.ResizeZ, def.Xan2D,
			def.Yan2D, def.Zan2D, def.Cost, def.IsTradable, def.Stackable, def.InventoryModel, def.WearPos1,
			def.WearPos2, def.WearPos3, def.Members, def.Zoom2D, def.XOffset2D, def.YOffset2D, def.Ambient,
			def.Contrast, SliceTextStr(def.Options), SliceTextStr(def.InterfaceOptions), def.MaleModel0, def.MaleModel1,
			def.MaleModel2, def.MaleOffset, def.MaleHeadModel, def.MaleHeadModel2, def.FemaleModel0, def.FemaleModel1,
			def.FemaleModel2, def.FemaleOffset, def.FemaleHeadModel, def.FemaleHeadModel2, def.NotedID,
			def.NotedTemplate, def.Team, def.Weight, def.ShiftClickDropIndex, def.BoughtID, def.BoughtTemplateID,
			def.PlaceholderID, def.PlaceholderTemplateID, SliceTextInt(def.ColorFind), SliceTextInt(def.ColorReplace),
			MapToStr(def.Params), SliceTextInt(def.CountCo), SliceTextInt(def.CountObj), SliceTextInt(def.TextureFind),
			SliceTextInt(def.TextureReplace), def.Category)

		if err != nil {
			fmt.Printf("Error executing prepared statement for item %s : %s\n", itemFiles[i], err)
		}
	}

}

func insertNPCData(cachePath string, database *sql.DB) {
	// get all files in /npc_defs
	npcFiles := getFileNames(fmt.Sprintf("%s\\npc_defs", cachePath))

	// for each file loop through
	for i := range npcFiles {
		// if file is not empty: (checked with getFileNames)
		// read file
		fileBytes, err := os.ReadFile(fmt.Sprintf("%s\\npc_defs\\%s", cachePath, npcFiles[i]))
		if err != nil {
			fmt.Printf("Error reading npc file %s to bytes: %s\n", npcFiles[i], err)
		}
		// unmarshal into an NPCEntry "def"
		def := NPCEntry{}
		if err = json.Unmarshal(fileBytes, &def); err != nil {
			fmt.Printf("Error unmarshalling NPC file %s : %s\n", npcFiles[i], err)
		}

		// SQL time
		statement := "INSERT OR REPLACE INTO npcs (id, name, size, models, chathead_models, standing_animation,  idle_rotate_left_animation, idle_rotate_right_animation, walking_animation, rotate_left_animation, rotate_right_animation, run_animation, run_rotate_180_animation, run_rotate_left_animation, run_rotate_right_animation, crawl_animation, crawl_rotate_180_animation, crawl_rotate_left_animation, crawl_rotate_right_animation, actions, is_minimap_visible, combat_level, width_scale, height_scale, has_render_priority, ambient, contrast, head_icon_sprite_index, head_icon_archive_ids, rotation_speed, varbit_id, varp_index, is_interactable, rotation_flag, is_pet, configs, params, category, recolor_to_find, recolor_to_replace, retexture_to_find, retexture_to_replace, is_follower, low_priority_follower_ops) VALUES (?, ?, ?, ?, ?, ?,  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		preparedStatement, err := database.Prepare(statement)
		if err != nil {
			fmt.Printf("Error preparing NPC insertion statement for item %s : %s\n", npcFiles[i], err)
			log.Fatal(err)
		}

		// I don't love this
		_, err = preparedStatement.Exec(def.ID, def.Name, def.Size, SliceTextInt(def.Models),
			SliceTextInt(def.ChatheadModels), def.StandingAnimation, def.IdleRotateLeftAnimation,
			def.IdleRotateRightAnimation, def.WalkingAnimation, def.RotateLeftAnimation, def.RotateRightAnimation,
			def.RunAnimation, def.RunRotate180Animation, def.RunRotateLeftAnimation, def.RunRotateRightAnimation,
			def.CrawlAnimation, def.CrawlRotate180Animation, def.CrawlRotateLeftAnimation,
			def.CrawlRotateRightAnimation, SliceTextStr(def.Actions), def.IsMinimapVisible, def.CombatLevel,
			def.WidthScale, def.HeightScale, def.HasRenderPriority, def.Ambient, def.Contrast,
			SliceTextInt(def.HeadIconSpriteIndex), SliceTextInt(def.HeadIconArchiveIDs), def.RotationSpeed,
			def.VarbitID, def.VarpIndex, def.IsInteractable, def.RotationFlag, def.IsPet, SliceTextInt(def.Configs),
			MapToStr(def.Params), def.Category, SliceTextInt(def.RecolorToFind), SliceTextInt(def.RecolorToReplace),
			SliceTextInt(def.RetextureToFind), SliceTextInt(def.RetextureToReplace), def.IsFollower,
			def.LowPriorityFollowerOps)

		if err != nil {
			fmt.Printf("Error executing prepared statement for npc %s : %s\n", npcFiles[i], err)
		}
	}
}

func insertObjectData(cachePath string, database *sql.DB) {
	// get all files in /object_defs
	objectFiles := getFileNames(fmt.Sprintf("%s\\object_defs", cachePath))

	// for each file loop through
	for i := range objectFiles {
		// if file is not empty: (checked with getFileNames)
		// read file
		fileBytes, err := os.ReadFile(fmt.Sprintf("%s\\object_defs\\%s", cachePath, objectFiles[i]))
		if err != nil {
			fmt.Printf("Error reading object file %s to bytes: %s\n", objectFiles[i], err)
		}
		// unmarshal into an ObjectEntry "def"
		def := ObjectEntry{}
		if err = json.Unmarshal(fileBytes, &def); err != nil {
			fmt.Printf("Error unmarshalling object file %s : %s\n", objectFiles[i], err)
		}

		// SQL time
		statement := "INSERT OR REPLACE INTO objects (id, name, decor_displacement, is_hollow, object_models, object_types, map_area_id, size_x, size_y, offset_x, offset_y, offset_height, merge_normals, wall_or_door, animation_id, varbit_id, ambient, contrast, recolor_to_find,  recolor_to_replace, retexture_to_find, texture_to_replace, actions, interact_type, map_scene_id, blocking_mask, shadow, model_size_x, model_size_y, model_size_height, object_id, obstructs_ground, contoured_ground, supports_items, config_change_dest, category, is_rotated, varp_id, ambient_sound_id, ambient_sound_ids, ambient_sound_retain, ambient_sound_distance, ambient_sound_change_ticks_min, ambient_sound_change_ticks_max, params, a_bool_2111, blocks_projectile, randomize_anim_start) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?,  ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
		preparedStatement, err := database.Prepare(statement)
		if err != nil {
			fmt.Printf("Error preparing object insertion statement for item %s : %s\n", objectFiles[i], err)
			log.Fatal(err)
		}

		// I don't love this
		_, err = preparedStatement.Exec(def.ID, def.Name, def.DecorDisplacement, def.IsHollow,
			SliceTextInt(def.ObjectModels), SliceTextInt(def.ObjectTypes), def.MapAreaID, def.SizeX, def.SizeY,
			def.OffsetX, def.OffsetY, def.OffsetHeight, def.MergeNormals, def.WallOrDoor, def.AnimationID, def.VarbitID,
			def.Ambient, def.Contrast, SliceTextInt(def.RecolorToFind), SliceTextInt(def.RecolorToReplace),
			SliceTextInt(def.RetextureToFind), SliceTextInt(def.TextureToReplace), SliceTextStr(def.Actions),
			def.InteractType, def.MapSceneID, def.BlockingMask, def.Shadow, def.ModelSizeX, def.ModelSizeY,
			def.ModelSizeHeight, def.ObjectID, def.ObstructsGround, def.ContouredGround, def.SupportsItems,
			SliceTextInt(def.ConfigChangeDest), def.Category, def.IsRotated, def.VarpID, def.AmbientSoundID,
			SliceTextInt(def.AmbientSoundIDs), def.AmbientSoundRetain, def.AmbientSoundDistance,
			def.AmbientSoundChangeTicksMin, def.AmbientSoundChangeTicksMax, MapToStr(def.Params), def.ABool2111,
			def.BlocksProjectile, def.RandomizeAnimStart)

		if err != nil {
			fmt.Printf("Error executing prepared statement for object %s : %s\n", objectFiles[i], err)
		}
	}
}

func PopulateTables(cachePath string, database *sql.DB) {
	fmt.Printf("Inserting items at %s\n", time.Now().Format(time.DateTime))
	insertItemData(cachePath, database)
	fmt.Printf("Inserting NPCs at %s\n", time.Now().Format(time.DateTime))
	insertNPCData(cachePath, database)
	fmt.Printf("Inserting objects at %s\n", time.Now().Format(time.DateTime))
	insertObjectData(cachePath, database)
	fmt.Printf("Finished at %s\n", time.Now().Format(time.DateTime))
}

func main() {
	dbName := "cache.db"
	cacheVersion := "2024-05-15-rev221"
	pathToCacheDump := fmt.Sprintf("C:\\Users\\cdfis\\Downloads\\dump-%s\\dump", cacheVersion)

	db := initializeDB(dbName)

	PopulateTables(pathToCacheDump, db)
}
