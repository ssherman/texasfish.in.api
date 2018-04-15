package main

import (
  "time"
  "github.com/gin-gonic/gin"
)

type LakeRecord struct {
  Id     int64
  FishTypeId   int64
  FishTypeName string
  AgeCategory string
  MethodCategory string
  Weight string
  Length string
  Date time.Time
  Angler string
  BaitOrLure string
  CreatedAt time.Time
  UpdatedAt time.Time
}

const selectLakeRecordsQuery = `
SELECT records.id as id, fish_type_id, fish_types.name as fish_type_name, age_category, method_category,
weight, length, date, angler, bait_or_lure, 
records.created_at as created_at, records.updated_at as updated_at
FROM records 
LEFT JOIN fish_types on fish_types.id = records.fish_type_id 
WHERE lake_id = $1 
ORDER BY id, method_category, age_category, fish_type_name
`

func (e *Env) lakeRecordsList(c *gin.Context) {
  var records []LakeRecord
  lakeId := c.Param("id")
  rows, err := e.db.Query(selectLakeRecordsQuery, lakeId)

  if err != nil {
    c.JSON(500, gin.H{"error": err})
    return
  }

  for rows.Next() {
    var lakeRecord LakeRecord
    rows.Scan(
      &lakeRecord.Id,
      &lakeRecord.FishTypeId,
      &lakeRecord.FishTypeName,
      &lakeRecord.AgeCategory,
      &lakeRecord.MethodCategory,
      &lakeRecord.Weight,
      &lakeRecord.Length,
      &lakeRecord.Date,
      &lakeRecord.Angler,
      &lakeRecord.BaitOrLure,
      &lakeRecord.CreatedAt,
      &lakeRecord.UpdatedAt)

    records = append(records, lakeRecord)
  }

  c.JSON(200, gin.H{"lake_records": records})
}

