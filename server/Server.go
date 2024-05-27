/* Server.go
2024, cdfisher
----------------
Development server for hosting an OSRS cache data API.
Proof of concept for a possible larger future project.
*/

package main

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	_ "github.com/ncruces/go-sqlite3"
	_ "github.com/ncruces/go-sqlite3/driver"
	_ "github.com/ncruces/go-sqlite3/embed"
	"log"
	"net/http"
)

var db *sql.DB
var err error

func notFound(c *gin.Context, key string, route string) {
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": fmt.Sprintf("Parameter %s not found for route %s", key,
		route)})
}

func fetchItems(query string, itemID string, c *gin.Context) []ItemEntry {
	var output []ItemEntry

	dbRows, err := db.QueryContext(c, query, itemID)
	if err != nil {
		log.Fatal("Error encountered executing query for search ", itemID, " : ", err)
	}

	i := 0

	for dbRows.Next() {
		rowData := ItemEntry{}
		err = dbRows.Scan(&rowData.ID, &rowData.Name, &rowData.Examine, &rowData.ResizeX, &rowData.ResizeY,
			&rowData.ResizeZ, &rowData.Xan2D, &rowData.Yan2D, &rowData.Zan2D, &rowData.Cost, &rowData.IsTradable,
			&rowData.Stackable, &rowData.InventoryModel, &rowData.WearPos1, &rowData.WearPos2, &rowData.WearPos3,
			&rowData.Members, &rowData.Zoom2D, &rowData.XOffset2D, &rowData.YOffset2D, &rowData.Ambient,
			&rowData.Contrast, &rowData.Options, &rowData.InterfaceOptions, &rowData.MaleModel0, &rowData.MaleModel1,
			&rowData.MaleModel2, &rowData.MaleOffset, &rowData.MaleHeadModel, &rowData.MaleHeadModel2,
			&rowData.FemaleModel0, &rowData.FemaleModel1, &rowData.FemaleModel2, &rowData.FemaleOffset,
			&rowData.FemaleHeadModel, &rowData.FemaleHeadModel2, &rowData.NotedID, &rowData.NotedTemplate,
			&rowData.Team, &rowData.Weight, &rowData.ShiftClickDropIndex, &rowData.BoughtID, &rowData.BoughtTemplateID,
			&rowData.PlaceholderID, &rowData.PlaceholderTemplateID, &rowData.ColorFind, &rowData.ColorReplace,
			&rowData.Params, &rowData.CountCo, &rowData.CountObj, &rowData.TextureFind, &rowData.TextureReplace,
			&rowData.Category)
		if err != nil {
			fmt.Println(err)
		}
		output = append(output, rowData)
		i++
	}
	return output
}

func BuildItemQuery(key string, c *gin.Context) string {
	queryType, ok := ItemQueryTypes[key]
	if !ok {
		notFound(c, key, "/items")
	}

	switch queryType {
	case 1:
		query := Queries[1]
		return fmt.Sprintf(query, "items", key)
	case 2:
		query := Queries[2]
		return fmt.Sprintf(query, "items", key)
	case 3:
		if c.Param(":value") == "null" {
			// needs exact match
			query := Queries[1]
			return fmt.Sprintf(query, "items", key)
		} else {
			// otherwise fuzzy match is good
			query := Queries[2]
			return fmt.Sprintf(query, "items", key)
		}
	default:
		// Unreachable case currently
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("The server somehow got " +
			"to an unreachable state.")})
	}
	return ""
}

func GetItems(c *gin.Context) {
	searchKey := c.Param("key")
	searchVal := c.Param("value")

	var results []ItemEntry

	queryString := BuildItemQuery(searchKey, c)
	results = append(results, fetchItems(queryString, searchVal, c)...)

	n := len(results)

	if n > 0 {
		fmt.Print(results, "\n")
		c.JSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No items matching query were found"})
	}
}

func BuildNPCQuery(key string, c *gin.Context) string {
	queryType, ok := NPCQueryTypes[key]
	if !ok {
		notFound(c, key, "/npcs")
	}

	switch queryType {
	case 1:
		query := Queries[1]
		return fmt.Sprintf(query, "npcs", key)
	case 2:
		query := Queries[2]
		return fmt.Sprintf(query, "npcs", key)
	case 3:
		if c.Param(":value") == "null" {
			// needs exact match
			query := Queries[1]
			return fmt.Sprintf(query, "npcs", key)
		} else {
			// otherwise fuzzy match is good
			query := Queries[2]
			return fmt.Sprintf(query, "npcs", key)
		}
	default:
		// Unreachable case currently
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("The server somehow got " +
			"to an unreachable state.")})
	}
	return ""
}

func fetchNPCs(query string, npcID string, c *gin.Context) []NPCEntry {
	var output []NPCEntry

	dbRows, err := db.QueryContext(c, query, npcID)
	if err != nil {
		log.Fatal("Error encountered executing query for search ", npcID, " : ", err)
	}

	i := 0

	for dbRows.Next() {
		rowData := NPCEntry{}
		err = dbRows.Scan(&rowData.ID, &rowData.Name, &rowData.Size, &rowData.Models,
			&rowData.ChatheadModels, &rowData.StandingAnimation, &rowData.IdleRotateLeftAnimation,
			&rowData.IdleRotateRightAnimation, &rowData.WalkingAnimation, &rowData.RotateLeftAnimation,
			&rowData.RotateRightAnimation, &rowData.RunAnimation, &rowData.RunRotate180Animation,
			&rowData.RunRotateLeftAnimation, &rowData.RunRotateRightAnimation, &rowData.CrawlAnimation,
			&rowData.CrawlRotate180Animation, &rowData.CrawlRotateLeftAnimation, &rowData.CrawlRotateRightAnimation,
			&rowData.Actions, &rowData.IsMinimapVisible, &rowData.CombatLevel, &rowData.WidthScale,
			&rowData.HeightScale, &rowData.HasRenderPriority, &rowData.Ambient, &rowData.Contrast,
			&rowData.HeadIconSpriteIndex, &rowData.HeadIconArchiveIDs, &rowData.RotationSpeed, &rowData.VarbitID,
			&rowData.VarpIndex, &rowData.IsInteractable, &rowData.RotationFlag, &rowData.IsPet, &rowData.Configs,
			&rowData.Params, &rowData.Category, &rowData.RecolorToFind, &rowData.RecolorToReplace,
			&rowData.RetextureToFind, &rowData.RetextureToReplace, &rowData.IsFollower, &rowData.LowPriorityFollowerOps)
		if err != nil {
			fmt.Println(err)
		}
		output = append(output, rowData)
		i++
	}
	return output
}

func GetNPCs(c *gin.Context) {
	searchKey := c.Param("key")
	searchVal := c.Param("value")

	var results []NPCEntry

	queryString := BuildNPCQuery(searchKey, c)
	results = append(results, fetchNPCs(queryString, searchVal, c)...)

	n := len(results)

	if n > 0 {
		fmt.Print(results, "\n")
		c.JSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No NPCs matching query were found"})
	}
}

func fetchObjects(query string, objectID string, c *gin.Context) []ObjectEntry {
	var output []ObjectEntry
	dbRows, err := db.QueryContext(c, query, objectID)
	if err != nil {
		log.Fatal("Error encountered executing query for search ", objectID, " : ", err)
	}

	i := 0

	for dbRows.Next() {
		rowData := ObjectEntry{}
		err = dbRows.Scan(&rowData.ID, &rowData.Name, &rowData.DecorDisplacement, &rowData.IsHollow,
			&rowData.ObjectModels, &rowData.ObjectTypes, &rowData.MapAreaID, &rowData.SizeX, &rowData.SizeY,
			&rowData.OffsetX, &rowData.OffsetY, &rowData.OffsetHeight, &rowData.MergeNormals, &rowData.WallOrDoor,
			&rowData.AnimationID, &rowData.VarbitID, &rowData.Ambient, &rowData.Contrast, &rowData.RecolorToFind,
			&rowData.RecolorToReplace, &rowData.RetextureToFind, &rowData.TextureToReplace, &rowData.Actions,
			&rowData.InteractType, &rowData.MapSceneID, &rowData.BlockingMask, &rowData.Shadow, &rowData.ModelSizeX,
			&rowData.ModelSizeY, &rowData.ModelSizeHeight, &rowData.ObjectID, &rowData.ObstructsGround,
			&rowData.ContouredGround, &rowData.SupportsItems, &rowData.ConfigChangeDest, &rowData.Category,
			&rowData.IsRotated, &rowData.VarpID, &rowData.AmbientSoundID, &rowData.AmbientSoundIDs,
			&rowData.AmbientSoundRetain, &rowData.AmbientSoundDistance, &rowData.AmbientSoundChangeTicksMin,
			&rowData.AmbientSoundChangeTicksMax, &rowData.Params, &rowData.ABool2111, &rowData.BlocksProjectile,
			&rowData.RandomizeAnimStart)
		if err != nil {
			fmt.Println(err)
		}
		output = append(output, rowData)
		i++
	}
	return output
}

func BuildObjectQuery(key string, c *gin.Context) string {
	queryType, ok := ObjectQueryTypes[key]
	if !ok {
		notFound(c, key, "/objects")
	}

	switch queryType {
	case 1:
		query := Queries[1]
		return fmt.Sprintf(query, "objects", key)
	case 2:
		query := Queries[2]
		return fmt.Sprintf(query, "objects", key)
	case 3:
		if c.Param(":value") == "null" {
			// needs exact match
			query := Queries[1]
			return fmt.Sprintf(query, "objects", key)
		} else {
			// otherwise fuzzy match is good
			query := Queries[2]
			return fmt.Sprintf(query, "objects", key)
		}
	default:
		// Unreachable case currently
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintf("The server somehow got " +
			"to an unreachable state.")})
	}
	return ""
}

func GetObjects(c *gin.Context) {
	searchKey := c.Param("key")
	searchVal := c.Param("value")

	var results []ObjectEntry

	queryString := BuildObjectQuery(searchKey, c)
	results = append(results, fetchObjects(queryString, searchVal, c)...)

	n := len(results)

	if n > 0 {
		fmt.Print(results, "\n")
		c.JSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "No objects matching query were found"})
	}
}

func initializeRouter() *gin.Engine {
	r := gin.Default()
	// TODO implement query params so these can be combined
	r.GET("items/:key/:value", GetItems)
	r.GET("npcs/:key/:value", GetNPCs)
	r.GET("objects/:key/:value", GetObjects)
	return r
}

func main() {
	db, err = sql.Open("sqlite3", "cache.db")
	if err != nil {
		log.Fatal("Failed to open database: ", err)
	}

	defer db.Close()

	router := initializeRouter()
	router.Run("localhost:8080")
}
