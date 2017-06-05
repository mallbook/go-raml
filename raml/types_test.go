package raml

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestTypeInType(t *testing.T) {
	apiDef := new(APIDefinition)
	Convey("Type in property's type", t, func() {
		err := ParseFile("./samples/types.raml", apiDef)
		So(err, ShouldBeNil)

		action, ok := apiDef.Types["Action"]
		So(ok, ShouldBeTrue)

		name := action.GetProperty("name")
		So(name.Type, ShouldEqual, "string")

		recurring := action.GetProperty("recurring")
		So(recurring.TypeString(), ShouldEqual, "Actionrecurring")

		// check the inline type
		ar, ok := apiDef.Types["Actionrecurring"]
		So(ok, ShouldBeTrue)

		// Must work via .GetPropert
		period := ar.GetProperty("period")
		So(period.TypeString(), ShouldEqual, "integer")

		// Also must work via .ToProperty
		var prop Property
		for k, p := range action.Properties {
			if k == "recurring" {
				prop = ToProperty(k, p)
				break
			}
		}
		So(prop.TypeString(), ShouldEqual, "Actionrecurring")
	})
}