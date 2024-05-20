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

func GetItemsFromID(c *gin.Context) {
	itemID := c.Param("id")

	var results []ItemEntry

	statement := "SELECT * FROM items WHERE id == ? ORDER BY id"
	dbRows, err := db.QueryContext(c, statement, itemID)
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
		results = append(results, rowData)
		i++
	}

	if i > 0 {
		fmt.Print(results, "\n")
		c.JSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "item id not found"})
	}
}

func GetNPCsFromID(c *gin.Context) {
	npcID := c.Param("id")

	var results []NPCEntry

	statement := "SELECT * FROM npcs WHERE id == ? ORDER BY id"
	dbRows, err := db.QueryContext(c, statement, npcID)
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
		results = append(results, rowData)
		i++
	}

	if i > 0 {
		fmt.Print(results, "\n")
		c.JSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "npc id not found"})
	}
}

func GetObjectsFromID(c *gin.Context) {
	objectID := c.Param("id")

	var results []ObjectEntry

	statement := "SELECT * FROM objects WHERE id == ? ORDER BY id"
	dbRows, err := db.QueryContext(c, statement, objectID)
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
		results = append(results, rowData)
		i++
	}

	if i > 0 {
		fmt.Print(results, "\n")
		c.JSON(http.StatusOK, results)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "object id not found"})
	}
}

func initializeRouter() *gin.Engine {
	r := gin.Default()
	r.GET("items/:id", GetItemsFromID)
	r.GET("npcs/:id", GetNPCsFromID)
	r.GET("objects/:id", GetObjectsFromID)
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
