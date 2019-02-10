package main

import (
  "time"
  "github.com/gin-gonic/gin"
)

type Event struct {
  Id     int64
  Title   string
  LakeId int64
  LakeName string
  Description string
  DateTimeStart time.Time
  DateTimeEnd time.Time
  TimeZone string
  Url string
  SubmittedById int64
  Owner string
  Approved bool
  CreatedAt time.Time
  UpdatedAt time.Time
}

const selectEventsQuery = `
SELECT events.id as id,
events.title as title,
events.lake_id as lake_id,
lakes.name as lake_name,
events.description as description,
datetime_start,
datetime_end,
time_zone,
events.url as url,
submitted_by_id,
owner,
approved,
events.created_at as created_at,
events.updated_at as updated_at
FROM events
INNER JOIN lakes on lakes.id = events.lake_id
WHERE datetime_start >= $1 AND datetime_start <= $2
ORDER BY datetime_start
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
      &event.LakeId,
      &event.LakeName,
      &event.Description,
      &event.DateTimeStart,
      &event.DateTimeEnd,
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
