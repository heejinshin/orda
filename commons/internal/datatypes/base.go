package datatypes

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"github.com/knowhunger/ortoo/commons/log"
	"github.com/knowhunger/ortoo/commons/model"
)

type baseDatatype struct {
	Key           string
	id            model.DUID
	opID          *model.OperationID
	TypeOf        model.TypeOfDatatype
	state         model.StateOfDatatype
	finalDatatype model.CommonDatatype
	Logger        *log.OrtooLog
}

// PublicBaseDatatypeInterface is a public interface for a datatype.
type PublicBaseDatatypeInterface interface {
	GetType() model.TypeOfDatatype
	GetState() model.StateOfDatatype
}

func newBaseDatatype(key string, t model.TypeOfDatatype, cuid model.CUID) (*baseDatatype, error) {
	duid, err := model.NewDUID()
	if err != nil {
		return nil, log.OrtooErrorf(err, "fail to create base datatype due to duid")
	}
	return &baseDatatype{
		Key:    key,
		id:     duid,
		TypeOf: t,
		opID:   model.NewOperationIDWithCuid(cuid),
		state:  model.StateOfDatatype_DUE_TO_CREATE,
		Logger: log.NewOrtooLogWithTag(fmt.Sprintf("%s", duid)[:8]),
	}, nil
}

func (b *baseDatatype) GetCUID() string {
	return hex.EncodeToString(b.opID.CUID)
}

func (b *baseDatatype) GetEra() uint32 {
	return b.opID.GetEra()
}

func (b *baseDatatype) String() string {
	return fmt.Sprintf("%s", b.id)
}

func (b *baseDatatype) executeLocalBase(op model.Operation) (interface{}, error) {
	b.SetNextOpID(op)
	return op.ExecuteLocal(b.finalDatatype)
}

func (b *baseDatatype) Replay(op model.Operation) error {
	if bytes.Compare(b.opID.CUID, op.GetBase().ID.CUID) == 0 {
		_, err := b.executeLocalBase(op)
		if err != nil {
			return log.OrtooErrorf(err, "fail to replay local operation")
		}
	} else {
		b.executeRemoteBase(op)
	}
	return nil
}

func (b *baseDatatype) SetNextOpID(op model.Operation) {
	op.GetBase().SetOperationID(b.opID.Next())
}

func (b *baseDatatype) executeRemoteBase(op model.Operation) {
	op.ExecuteRemote(b.finalDatatype)
}

func (b *baseDatatype) GetType() model.TypeOfDatatype {
	return b.TypeOf
}

func (b *baseDatatype) GetState() model.StateOfDatatype {
	return b.state
}

func (b *baseDatatype) GetDatatype() model.StateOfDatatype {
	return b.state
}

func (b *baseDatatype) SetFinalDatatype(finalDatatype model.CommonDatatype) {
	b.finalDatatype = finalDatatype
}

func (b *baseDatatype) GetFinalDatatype() model.CommonDatatype {
	return b.finalDatatype
}

func (b *baseDatatype) SetOpID(opID *model.OperationID) {
	b.opID = opID
}

func (b *baseDatatype) GetKey() string {
	return b.Key
}

func (b *baseDatatype) GetDUID() model.DUID {
	return b.id
}

func (b *baseDatatype) SetDUID(duid model.DUID) {
	b.id = duid
}

func (b *baseDatatype) SetState(state model.StateOfDatatype) {
	b.state = state
}
