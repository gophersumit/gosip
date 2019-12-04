package api

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestItems(t *testing.T) {
	checkClient(t)

	web := NewSP(spClient).Web()
	newListTitle := uuid.New().String()
	if _, err := web.Lists().Add(newListTitle, nil); err != nil {
		t.Error(err)
	}
	list := web.Lists().GetByTitle(newListTitle)
	entType, err := list.GetEntityType()
	if err != nil {
		t.Error(err)
	}
	startedOn := time.Now()

	t.Run("AddWithoutMetadataType", func(t *testing.T) {
		body := []byte(`{"Title":"Item"}`)
		if _, err := list.Items().Add(body); err != nil {
			t.Error(err)
		}
	})

	t.Run("AddResponse", func(t *testing.T) {
		body := []byte(`{"Title":"Item"}`)
		item, err := list.Items().Add(body)
		if err != nil {
			t.Error(err)
		}
		if item.Data().ID == 0 {
			t.Error("can't get item properly")
		}
	})

	t.Run("AddSeries", func(t *testing.T) {
		for i := 1; i < 10; i++ {
			metadata := make(map[string]interface{})
			metadata["__metadata"] = map[string]string{"type": entType}
			metadata["Title"] = fmt.Sprintf("Item %d", i)
			body, _ := json.Marshal(metadata)
			if _, err := list.Items().Add(body); err != nil {
				t.Error(err)
			}
		}
	})

	t.Run("Get", func(t *testing.T) {
		items, err := list.Items().Top(100).OrderBy("Title", false).Get()
		if err != nil {
			t.Error(err)
		}
		if len(items.Data()) == 0 {
			t.Error("can't get items properly")
		}
		if items.Data()[0].Data().ID == 0 {
			t.Error("can't get items properly")
		}
	})

	t.Run("GetByID", func(t *testing.T) {
		item, err := list.Items().GetByID(1).Get()
		if err != nil {
			t.Error(err)
		}
		if item.Data().ID == 0 {
			t.Error("can't get item properly")
		}
	})

	t.Run("Get/Unmarshal", func(t *testing.T) {
		item, err := list.Items().GetByID(1).Get()
		if err != nil {
			t.Error(err)
		}
		if item.Data().ID == 0 {
			t.Error("can't get item ID property properly")
		}
		if item.Data().Title == "" {
			t.Error("can't get item Title property properly")
		}
		if item.Data().Created.Day() != startedOn.Day() {
			t.Error("can't get item Created property properly")
		}
		if item.Data().Modified.Day() != startedOn.Day() {
			t.Error("can't get item Modified property properly")
		}
	})

	t.Run("Get/CustomUnmarshal", func(t *testing.T) {
		items, err := list.Items().Select("Id,RoleAssignments").Expand("RoleAssignments").Top(5).Get()
		if err != nil {
			t.Error(err)
		}
		obj := []*struct {
			ID              int                      `json:"Id"`
			RoleAssignments []map[string]interface{} `json:"RoleAssignments"`
		}{}
		if err := items.Unmarshal(&obj); err != nil {
			t.Error(err)
		}
		if len(obj[0].RoleAssignments) == 0 {
			t.Error("can't parse response object")
		}
	})

	t.Run("GetByCAML", func(t *testing.T) {
		caml := `
			<View>
				<Query>
					<Where>
						<Eq>
							<FieldRef Name='ID' />
							<Value Type='Number'>3</Value>
						</Eq>
					</Where>
				</Query>
			</View>
		`
		data, err := list.Items().GetByCAML(caml)
		if err != nil {
			t.Error(err)
		}
		if len(data.Data()) != 1 {
			t.Error("incorrect number of items")
		}
		if data.Data()[0].Data().ID != 3 {
			t.Error("incorrect response")
		}
	})

	if _, err := list.Delete(); err != nil {
		t.Error(err)
	}

}
