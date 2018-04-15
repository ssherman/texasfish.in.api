package main

import (
  "time"
  "github.com/gin-gonic/gin"
)

type FishType struct {
  Id     int64
  Name string
  Slug string
  RawName string
  Description string
  CreatedAt time.Time
  UpdatedAt time.Time
}

const selectFishTypesQuery = `
SELECT id, name, slug, raw_name, description, created_at, updated_at
FROM Fish_types
ORDER BY name
`

func (e *Env) fishTypesList(c *gin.Context) {
  var fishTypes []FishType
  rows, err := e.db.Query(selectFishTypesQuery)

  if err != nil {
    c.JSON(500, gin.H{"error": err})
    return
  }

  for rows.Next() {
    var fishType FishType
    rows.Scan(
      &fishType.Id,
      &fishType.Name,
      &fishType.Slug,
      &fishType.RawName,
      &fishType.Description,
      &fishType.CreatedAt,
      &fishType.UpdatedAt)
    fishTypes = append(fishTypes, fishType)
  }
  
  c.JSON(200, gin.H{"fish_types": fishTypes})
}
