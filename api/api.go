// Package api implement API for communication with Neo4j
package api

import (
	"fmt"
	"log"
	"strconv"

	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
)

// Bacon struct wraps information related
// to Neo4j communication access and results
// from queries
type Bacon struct {
	conn          bolt.Conn
	url           string
	oldlistquery  string
	newlistquery  string
	oldlistresult [][]interface{}
	newlistresult [][]interface{}
	result        []*Restaurant
}

// Get ask Neo4j for a new restaurant list
func Get() ([]*Restaurant, error) {
	c, err := NewClient("resources/conf/neo4j_local.yml")
	if err != nil {
		log.Fatal("failed to get neogo client: ", err)
	}

	// Open Neo4j connection
	if err := c.Connection(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := c.CloseConn(); err != nil {
			log.Fatal("failed to close neo4go connection: ", err)
		}
	}()

	b := newBacon(c)

	// Get lists
	if err := b.getRestaurants(); err != nil {
		return nil, fmt.Errorf("failed to get restaurants: %s", err)
	}
	b.mapResult()
	b.conn.Close()
	return b.result, nil
}

// newBacon retuns a new Bacon instance
func newBacon(c Client) *Bacon {
	return &Bacon{
		conn:         c.Conn,
		url:          c.url,
		oldlistquery: "MATCH (r:Restaurant)-->(b:Bacon) where b.last_points is not null return r.id as Rid, r.name as Restaurant ORDER BY b.last_points DESC",
		newlistquery: "MATCH (r:Restaurant)-->(b:Bacon) where b.points is not null return r.id as Rid, r.name as Restaurant ORDER BY b.points DESC",
	}
}

// getRestaurants execute Neo4j query
// to retrive data and generate new ranking
func (b *Bacon) getRestaurants() error {

	// Query the old list of restaurants
	oldlist, err := b.queryNeo4j(b.oldlistquery)
	if err != nil {
		return fmt.Errorf("failed to receive the old list %s", err)
	}
	b.oldlistresult = oldlist

	// Query the new list of restaurants
	newlist, err := b.queryNeo4j(b.newlistquery)
	if err != nil {
		return fmt.Errorf("failed to receive the new list %s", err)
	}
	b.newlistresult = newlist
	return nil
}

// Query Neo4j
func (b *Bacon) queryNeo4j(query string) ([][]interface{}, error) {
	data, _, _, err := b.conn.QueryNeoAll(query, nil)
	if err != nil {
		return nil, err
	}
	return data, nil
}

// mapResults iterates over the lists of restaurants
// and creates maps containing the restaurant name
// and the position occupied by the restaurant in the
// original list
func (b *Bacon) mapResult() {

	var restaurants []*Restaurant

	for i, value := range b.newlistresult {
		r := &Restaurant{
			ID:     value[0].(string),
			Name:   value[1].(string),
			Newpos: int32(i) + 1,
		}
		restaurants = append(restaurants, r)
	}

	for _, r := range restaurants {
		r.Change = "N/A"
		for i, value := range b.oldlistresult {
			if r.ID == value[0].(string) {
				r.Lastpos = int32(i) + 1
				strChange := strconv.Itoa(int(r.Lastpos - r.Newpos))
				r.Change = strChange
			}
		}
	}
	b.result = restaurants
}
