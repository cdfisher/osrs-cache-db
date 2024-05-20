/* ResponseEntries.go
2024, cdfisher
----------------
Structs for  for marshalling OSRS cache DB query results into JSON.
These structs use different field types in places than used in Entries.go, they are not
interchangeable.

Includes all keys found in /item_defs/, /npc_defs/, and /object_defs/ as of
cache 221.7 (2024-05-15-rev221)
*/

/*
// TODO
ResponseEntries TODOs:
----------------
- Arrays of ints are currently being marshalled into strings. Marshal them into arrays of ints
- Marshal string arrays into arrays of strings instead of single strings
- Add handling for unknown keys
- Add support for other objects in cache: dbtables, param_defs, ???
- Update Params fields to use generics or reflection rather than interface{}
- Reorder fields to match appearance in defs

*/

package main

type ItemEntry struct {
	ID                    int    `json:"id"`
	Name                  string `json:"name"`
	Examine               string `json:"examine"`
	ResizeX               int    `json:"resizeX"`
	ResizeY               int    `json:"resizeY"`
	ResizeZ               int    `json:"resizeZ"`
	Xan2D                 int    `json:"xan2D"`
	Yan2D                 int    `json:"yan2D"`
	Zan2D                 int    `json:"zan2D"`
	Cost                  int    `json:"cost"`
	IsTradable            bool   `json:"isTradable"`
	Stackable             int    `json:"stackable"`
	InventoryModel        int    `json:"inventoryModel"`
	WearPos1              int    `json:"wearPos1"`
	WearPos2              int    `json:"wearPos2"`
	WearPos3              int    `json:"wearPos3"`
	Members               bool   `json:"members"`
	Zoom2D                int    `json:"zoom2D"`
	XOffset2D             int    `json:"xOffset2d"`
	YOffset2D             int    `json:"yOffset2d"`
	Ambient               int    `json:"ambient"`
	Contrast              int    `json:"contrast"`
	Options               string `json:"options"`
	InterfaceOptions      string `json:"interfaceOptions"`
	MaleModel0            int    `json:"maleModel0"`
	MaleModel1            int    `json:"maleModel1"`
	MaleModel2            int    `json:"maleModel2"`
	MaleOffset            int    `json:"maleOffset"`
	MaleHeadModel         int    `json:"maleHeadModel"`
	MaleHeadModel2        int    `json:"maleHeadModel2"`
	FemaleModel0          int    `json:"femaleModel0"`
	FemaleModel1          int    `json:"femaleModel1"`
	FemaleModel2          int    `json:"femaleModel2"`
	FemaleOffset          int    `json:"femaleOffset"`
	FemaleHeadModel       int    `json:"femaleHeadModel"`
	FemaleHeadModel2      int    `json:"femaleHeadModel2"`
	NotedID               int    `json:"notedID"`
	NotedTemplate         int    `json:"notedTemplate"`
	Team                  int    `json:"team"`
	Weight                int    `json:"weight"`
	ShiftClickDropIndex   int    `json:"shiftClickDropIndex"`
	BoughtID              int    `json:"boughtId"`
	BoughtTemplateID      int    `json:"boughtTemplateId"`
	PlaceholderID         int    `json:"placeholderId"`
	PlaceholderTemplateID int    `json:"placeholderTemplateId"`
	ColorFind             string `json:"colorFind"`
	ColorReplace          string `json:"colorReplace"`
	Params                string `json:"params"`
	CountCo               string `json:"countCo"`
	CountObj              string `json:"countObj"`
	TextureFind           string `json:"textureFind"`
	TextureReplace        string `json:"textureReplace"`
	Category              int    `json:"category"`
}

type NPCEntry struct {
	ID                        int    `json:"id"`
	Name                      string `json:"name"`
	Size                      int    `json:"size"`
	Models                    string `json:"models"`
	ChatheadModels            string `json:"chatheadModels"`
	StandingAnimation         int    `json:"standingAnimation"`
	IdleRotateLeftAnimation   int    `json:"idleRotateLeftAnimation"`
	IdleRotateRightAnimation  int    `json:"idleRotateRightAnimation"`
	WalkingAnimation          int    `json:"walkingAnimation"`
	RotateLeftAnimation       int    `json:"rotateLeftAnimation"`
	RotateRightAnimation      int    `json:"rotateRightAnimation"`
	RunAnimation              int    `json:"runAnimation"`
	RunRotate180Animation     int    `json:"runRotate180Animation"`
	RunRotateLeftAnimation    int    `json:"runRotateLeftAnimation"`
	RunRotateRightAnimation   int    `json:"runRotateRightAnimation"`
	CrawlAnimation            int    `json:"crawlAnimation"`
	CrawlRotate180Animation   int    `json:"crawlRotate180Animation"`
	CrawlRotateLeftAnimation  int    `json:"crawlRotateLeftAnimation"`
	CrawlRotateRightAnimation int    `json:"crawlRotateRightAnimation"`
	Actions                   string `json:"actions"`
	IsMinimapVisible          bool   `json:"isMinimapVisible"`
	CombatLevel               int    `json:"combatLevel"`
	WidthScale                int    `json:"widthScale"`
	HeightScale               int    `json:"heightScale"`
	HasRenderPriority         bool   `json:"hasRenderPriority"`
	Ambient                   int    `json:"ambient"`
	Contrast                  int    `json:"contrast"`
	HeadIconSpriteIndex       string `json:"headIconSpriteIndex"`
	HeadIconArchiveIDs        string `json:"headIconArchiveIds"`
	RotationSpeed             int    `json:"rotationSpeed"`
	VarbitID                  int    `json:"varbitId"`
	VarpIndex                 int    `json:"varpIndex"`
	IsInteractable            bool   `json:"isInteractable"`
	RotationFlag              bool   `json:"rotationFlag"`
	IsPet                     bool   `json:"isPet"`
	Configs                   string `json:"configs"`
	Params                    string `json:"params"`
	Category                  int    `json:"category"`
	RecolorToFind             string `json:"recolorToFind"`
	RecolorToReplace          string `json:"recolorToReplace"`
	RetextureToFind           string `json:"retextureToFind"`
	RetextureToReplace        string `json:"retextureToReplace"`
	IsFollower                bool   `json:"isFollower"`
	LowPriorityFollowerOps    bool   `json:"lowPriorityFollowerOps"`
}

type ObjectEntry struct {
	ID                         int    `json:"id"`
	Name                       string `json:"name"`
	DecorDisplacement          int    `json:"decorDisplacement"`
	IsHollow                   bool   `json:"isHollow"`
	ObjectModels               string `json:"objectModels"`
	ObjectTypes                string `json:"objectTypes"`
	MapAreaID                  int    `json:"mapAreaId"`
	SizeX                      int    `json:"sizeX"`
	SizeY                      int    `json:"sizeY"`
	OffsetX                    int    `json:"offsetX"`
	OffsetY                    int    `json:"offsetY"`
	OffsetHeight               int    `json:"offsetHeight"`
	MergeNormals               bool   `json:"mergeNormals"`
	WallOrDoor                 int    `json:"wallOrDoor"`
	AnimationID                int    `json:"animationID"`
	VarbitID                   int    `json:"varbitID"`
	Ambient                    int    `json:"ambient"`
	Contrast                   int    `json:"contrast"`
	RecolorToFind              string `json:"recolorToFind"`
	RecolorToReplace           string `json:"recolorToReplace"`
	RetextureToFind            string `json:"retextureToFind"`
	TextureToReplace           string `json:"textureToReplace"`
	Actions                    string `json:"actions"`
	InteractType               int    `json:"interactType"`
	MapSceneID                 int    `json:"mapSceneID"`
	BlockingMask               int    `json:"blockingMask"`
	Shadow                     bool   `json:"shadow"`
	ModelSizeX                 int    `json:"modelSizeX"`
	ModelSizeY                 int    `json:"modelSizeY"`
	ModelSizeHeight            int    `json:"modelSizeHeight"`
	ObjectID                   int    `json:"objectID"`
	ObstructsGround            bool   `json:"obstructsGround"`
	ContouredGround            int    `json:"contouredGround"`
	SupportsItems              int    `json:"supportsItems"`
	ConfigChangeDest           string `json:"configChangeDest"`
	Category                   int    `json:"category"`
	IsRotated                  bool   `json:"isRotated"`
	VarpID                     int    `json:"varpID"`
	AmbientSoundID             int    `json:"ambientSoundId"`
	AmbientSoundIDs            string `json:"ambientSoundIds"`
	AmbientSoundRetain         int    `json:"ambientSoundRetain"`
	AmbientSoundDistance       int    `json:"ambientSoundDistance"`
	AmbientSoundChangeTicksMin int    `json:"ambientSoundChangeTicksMin"`
	AmbientSoundChangeTicksMax int    `json:"ambientSoundChangeTicksMax"`
	Params                     string `json:"params"`
	ABool2111                  bool   `json:"aBool2111"`
	BlocksProjectile           bool   `json:"blocksProjectile"`
	RandomizeAnimStart         bool   `json:"randomizeAnimStart"`
}
