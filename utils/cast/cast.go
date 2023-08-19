package cast

import (
	"github.com/google/uuid"
)

func CastStringToUUID(id string)( *uuid.UUID,error){
	new_id,err := uuid.Parse(id)
	if err != nil {
		return nil,err
	}
	return &new_id,nil
}