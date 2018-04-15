package main

import (
  "database/sql"
  "github.com/lib/pq"
  "fmt"
  "time"
  "github.com/gin-gonic/gin"
)

type Lake struct {
  Id     int64
  Name   string
  DetailsUri string
  SurfaceAreaInAcres int64
  MaxDepthInFeet int64
  YearImpounded int64
  ConservationPoolElevation float32
  PercentageFull float32
  Fluctuation string
  NormalClarity string
  WaterDataUri string
  ReservoirControllingAuthority string
  AquaticVegetation string
  PredominantFishSpecies []string
  AnglingOpportunitiesDescription string
  AnglingOpportunitiesDetails PropertyMap
  FishingRegulations string
  LakeMaps string
  LatestSurveyReport string
  StructureAndCoverDescription string
  TipsAndTactics string
  FishingRecordsUri string
  StockingHistoryUri string
  LocationDesc string
  CreatedAt time.Time
  UpdatedAt time.Time
}

const selectLakesQuery = `
SELECT id,
name,
details_uri,
surface_area_in_acres,
max_depth_in_feet,
year_impounded,
conservation_pool_elevation_in_ft_msl,
percentage_full,
fluctuation,
normal_clarity,
water_data_uri,
reservoir_controlling_authority,
aquatic_vegetation,
predominant_fish_species,
angling_opportunities_description,
angling_opportunities_details,
fishing_regulations,
lake_maps,
latest_survey_report,
structure_and_cover_description,
tips_and_tactics,
fishing_records_uri,
stocking_history_uri,
location_desc,
created_at,
updated_at
FROM lakes
ORDER BY name
`

const selectLakeQuery = `
SELECT id,
name,
details_uri,
surface_area_in_acres,
max_depth_in_feet,
year_impounded,
conservation_pool_elevation_in_ft_msl,
percentage_full,
fluctuation,
normal_clarity,
water_data_uri,
reservoir_controlling_authority,
aquatic_vegetation,
predominant_fish_species,
angling_opportunities_description,
angling_opportunities_details,
fishing_regulations,
lake_maps,
latest_survey_report,
structure_and_cover_description,
tips_and_tactics,
fishing_records_uri,
stocking_history_uri,
location_desc,
created_at,
updated_at
FROM lakes WHERE id = $1
`

func lakeShow(c *gin.Context) {
  dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
      "SSherman", "texasfish_in_development")
  db, err := sql.Open("postgres", dbinfo)
  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
  }
  var lake Lake
  lakeId := c.Param("id")
  row := db.QueryRow(selectLakeQuery, lakeId)
  row.Scan(
    &lake.Id,
    &lake.Name,
    &lake.DetailsUri,
    &lake.SurfaceAreaInAcres,
    &lake.MaxDepthInFeet,
    &lake.YearImpounded,
    &lake.ConservationPoolElevation,
    &lake.PercentageFull,
    &lake.Fluctuation,
    &lake.NormalClarity,
    &lake.WaterDataUri,
    &lake.ReservoirControllingAuthority,
    &lake.AquaticVegetation,
    pq.Array(&lake.PredominantFishSpecies),
    &lake.AnglingOpportunitiesDescription,
    &lake.AnglingOpportunitiesDetails,
    &lake.FishingRegulations,
    &lake.LakeMaps,
    &lake.LatestSurveyReport,
    &lake.StructureAndCoverDescription,
    &lake.TipsAndTactics,
    &lake.FishingRecordsUri,
    &lake.StockingHistoryUri,
    &lake.LocationDesc,
    &lake.CreatedAt,
    &lake.UpdatedAt)

  c.JSON(200, gin.H{"lake": lake})

}

func lakes(c *gin.Context) {
  dbinfo := fmt.Sprintf("user=%s dbname=%s sslmode=disable",
      "SSherman", "texasfish_in_development")
  db, err := sql.Open("postgres", dbinfo)
  if err != nil {
    fmt.Println("error")
    fmt.Println(err)
  }
  var lakes []Lake
  rows, err := db.Query(selectLakesQuery)
  for rows.Next() {
    var lake Lake
    rows.Scan(
      &lake.Id,
      &lake.Name,
      &lake.DetailsUri,
      &lake.SurfaceAreaInAcres,
      &lake.MaxDepthInFeet,
      &lake.YearImpounded,
      &lake.ConservationPoolElevation,
      &lake.PercentageFull,
      &lake.Fluctuation,
      &lake.NormalClarity,
      &lake.WaterDataUri,
      &lake.ReservoirControllingAuthority,
      &lake.AquaticVegetation,
      pq.Array(&lake.PredominantFishSpecies),
      &lake.AnglingOpportunitiesDescription,
      &lake.AnglingOpportunitiesDetails,
      &lake.FishingRegulations,
      &lake.LakeMaps,
      &lake.LatestSurveyReport,
      &lake.StructureAndCoverDescription,
      &lake.TipsAndTactics,
      &lake.FishingRecordsUri,
      &lake.StockingHistoryUri,
      &lake.LocationDesc,
      &lake.CreatedAt,
      &lake.UpdatedAt)
    lakes = append(lakes, lake)
  }
  
  c.JSON(200, gin.H{"lakes": lakes})
}
