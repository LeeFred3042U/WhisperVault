package storage
import (

	"encoding/json"
	"os"
	"ModularVault/internal/models"
)

const FileName = "contacts.json"

func LoadContacts() []models.Contact {

	data, err := os.ReadFile(FileName)

	if err != nil {
		return []models.Contact{}
	}

	var contacts []models.Contact
	json.Unmarshal(data, &contacts)

	return contacts 
}


func SaveContacts(contacts []models.Contact){

	data, _ := json.MarshalIndent(contacts, "",  " ")

	os.WriteFile(FileName, data, 0644)
}

