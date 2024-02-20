package main

import (
	"encoding/json"
	"os"

	"github.com/supabase-community/postgrest-go"
)

var supabaseUrl string = os.Getenv("SUPABASE_URL")
var supabaseKey string = os.Getenv("SUPABASE_KEY")

// Initialize a new Supabase client
var client = postgrest.NewClient(supabaseUrl, "public", map[string]string{
	"apikey":        supabaseKey,
	"Authorization": "Bearer " + supabaseKey,
})

// Example function to fetch data from a Supabase table
func FetchDataFromSupabase() ([]map[string]interface{}, error) {
	res, _, err := client.From("your_table_name").Select("*", "", false).Execute()
	if err != nil {
		return nil, err
	}
	var data []map[string]interface{}
	err = json.Unmarshal(res, &data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

// Example function to insert data into a Supabase table
func InsertDataToSupabase(data map[string]interface{}) error {
	// Assuming the boolean is for upsert and the strings are for returning, duplicates, and upsert conditions which we'll leave as defaults for this example.
	_, _, err := client.From("your_table_name").Insert(data, false, "", "", "").Execute()
	if err != nil {
		return err
	}
	return nil
}
