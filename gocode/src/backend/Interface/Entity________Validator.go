package Interface

import (
	"gopkg.in/go-playground/validator.v9"
	"gopkg.in/mgo.v2/bson"
	"strings"
	"log"
)

//__ VALIDATE NAME USED ______________________________________________________//
func ValidateNameUsed(name validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.ValidateNameUsed\n")

	c := OpenSession(docname)
	n := strings.ToLower(name.Field().String())
	var result Name
	c.Find(bson.M{ID_Name: n}).One(&result)
	return result.Name == ""
}

//__ VALIDATE NAME EXIST _____________________________________________________//
func ValidateNameExist(name validator.FieldLevel) bool {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.ValidateNameExist()\n")

	c := OpenSession(docname)
	n := strings.ToLower(name.Field().String())
	var result Name
	c.Find(bson.M{ID_Name: n}).One(&result)
	return result.Name != ""
}

//__ VALIDATE UNIT USED ______________________________________________________//
func ValidateUnitUsed(unit Unit) bool {
/*________________________________________TESTING FUNCTION________________________________________*/
log.Println("\n\nInterface.ValidateUnitUsed()\n")

	c := OpenSession(docname)
	unit.Username = strings.ToLower(unit.Username)
	unit.Name = strings.ToLower(unit.Name)
	c.Find(bson.M{ID_Username: unit.Username, ID_Name: unit.Name}).One(&unit)
	return unit.Username == ""
}
