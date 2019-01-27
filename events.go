package main

import (
  "time"
  "github.com/gin-gonic/gin"
)

type Event struct {
  Id     int64
  Title   string
  Description string
  LakeId int64
  Date time.Time
  TimeStart time.Time
  TimeEnd time.Time
  TimeZone string
  Url string
  SubmittedById int64
  Owner string
  Approved bool
  CreatedAt time.Time
  UpdatedAt time.Time
}

const selectEventsQuery = `
SELECT id,
title,
description,
lake_id,
date,
time_start,
time_end,
time_zone,
url,
submitted_by_id,
owner,
approved,
created_at,
updated_at
FROM events
WHERE date >= $1 AND date <= $2
ORDER BY date, time_start
`

func (e *Env) eventsList(c *gin.Context) {
  var events []Event
  start := c.Query("start")
  end := c.Query("end")
  rows, err := e.db.Query(selectEventsQuery, start, end)

  if err != nil {
    c.JSON(500, gin.H{"error": err})
    return
  }

  for rows.Next() {
    var event Event
    rows.Scan(
      &event.Id,
      &event.Title,
      &event.Description,
      &event.LakeId,
      &event.Date,
      &event.TimeStart,
      &event.TimeEnd,
      &event.TimeZone,
      &event.Url,
      &event.SubmittedById,
      &event.Owner,
      &event.Approved,
      &event.CreatedAt,
      &event.UpdatedAt)
    events = append(events, event)
  }
  
  c.JSON(200, gin.H{"events": events})
}
