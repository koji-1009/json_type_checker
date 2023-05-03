package checker

import "testing"

// test ReadJsonFile function
func TestReadJsonFile(t *testing.T) {
	// test file path
	const path = "../fixture/users.json"
	// read json file
	json, err := ReadJsonFile(path)
	// check error
	if err != nil {
		t.Errorf("Error: %s", err)
	}

	// check json
	if json == nil {
		t.Errorf("Error: json is nil")
	}

	// check users
	if json["users"] == nil {
		t.Errorf("Error: users is nil")
	}

	// check users type
	if _, ok := json["users"].([]interface{}); !ok {
		t.Errorf("Error: users is not array")
	}

	// check users length
	if len(json["users"].([]interface{})) != 2 {
		t.Errorf("Error: users length is not 2")
	}

	// check first user
	if elliot := json["users"].([]interface{})[0].(map[string]interface{}); elliot == nil {
		t.Errorf("Error: first user is nil")
	} else {
		// name is Elliot, type is Reader, age is 23, social is map
		if elliot["name"] != "Elliot" {
			t.Errorf("Error: user name is not Elliot")
		}
		if elliot["type"] != "Reader" {
			t.Errorf("Error: user type is not Reader")
		}
		if elliot["age"].(float64) != 23 {
			t.Errorf("Error: user age is not 23")
		}
		if elliot["social"] == nil {
			t.Errorf("Error: user social is nil")
		} else {
			if social := elliot["social"].(map[string]interface{}); social == nil {
				t.Errorf("Error: user social is not map")
			} else {
				if social["facebook"] != "https://example.com/facebook/elliot" && social["twitter"] != "https://example.com/twitter/elliot" {
					t.Errorf("Error: user social is not map")
				}
			}
		}
	}

	// check second user
	if fraser := json["users"].([]interface{})[1].(map[string]interface{}); fraser == nil {
		t.Errorf("Error: first user is nil")
	} else {
		// name is Fraser, type is Author, age is 17, social is map
		if fraser["name"] != "Fraser" {
			t.Errorf("Error: user name is not Fraser")
		}
		if fraser["type"] != "Author" {
			t.Errorf("Error: user type is not Author")
		}
		if fraser["age"].(float64) != 17 {
			t.Errorf("Error: user age is not 17")
		}
		if fraser["social"] == nil {
			t.Errorf("Error: user social is nil")
		} else {
			if social := fraser["social"].(map[string]interface{}); social == nil {
				t.Errorf("Error: user social is not map")
			} else {
				if social["facebook"] != "https://example.com/facebook/fraser" && social["twitter"] != "https://example.com/twitter/fraser" {
					t.Errorf("Error: user social is not map")
				}
			}
		}
	}
}
