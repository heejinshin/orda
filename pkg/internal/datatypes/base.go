package datatypes

import (
	"encoding/json"
	"fmt"
	"github.com/orda-io/orda/pkg/context"
	"github.com/orda-io/orda/pkg/errors"
	"github.com/orda-io/orda/pkg/iface"
	"github.com/orda-io/orda/pkg/log"
	"github.com/orda-io/orda/pkg/model"
	"github.com/orda-io/orda/pkg/types"
	"github.com/orda-io/orda/pkg/utils"
)

// BaseDatatype is the base datatype which contains
type BaseDatatype struct {
	Key    string
	id     string
	opID   *model.OperationID
	TypeOf model.TypeOfDatatype
	state  model.StateOfDatatype
	ctx    *context.DatatypeContext
	iface.Datatype
}

// NewBaseDatatype creates a new base datatype
func NewBaseDatatype(
	key string,
	t model.TypeOfDatatype,
	clientCtx *context.ClientContext,
	state model.StateOfDatatype,
) *BaseDatatype {
	duid := types.NewUID()
	base := &BaseDatatype{
		Key:    key,
		id:     duid,
		TypeOf: t,
		opID:   model.NewOperationIDWithCUID(clientCtx.Client.CUID),
		state:  state,
	}
	base.ctx = context.NewDatatypeContext(clientCtx, base)
	return base
}

// GetCUID returns CUID of the client which this datatype subscribes to.
func (its *BaseDatatype) GetCUID() string {
	return its.opID.CUID
}

// GetEra returns the era of operation ID.
func (its *BaseDatatype) GetEra() uint32 {
	return its.opID.GetEra()
}

func (its *BaseDatatype) SetLogger(l *log.OrdaLog) {
	its.ctx.SetLogger(l)
}

func (its *BaseDatatype) String() string {
	return fmt.Sprintf("%s", its.id)
}

func (its *BaseDatatype) executeLocalBase(op iface.Operation) (interface{}, errors.OrdaError) {
	its.SetNextOpID(op)
	switch op.GetType() {
	case model.TypeOfOperation_TRANSACTION, model.TypeOfOperation_ERROR, model.TypeOfOperation_SNAPSHOT:
		return nil, nil
	}
	ret, err := its.ExecuteLocal(op)
	if err != nil {
		its.opID.RollBack()
	}
	return ret, err // should deliver err
}

func (its *BaseDatatype) executeRemoteBase(op iface.Operation) {
	its.opID.SyncLamport(op.GetID().Lamport)
	_, _ = its.ExecuteRemote(op)
}

// Replay replays an already executed operation.
func (its *BaseDatatype) Replay(op iface.Operation) errors.OrdaError {
	if its.opID.CUID == op.GetID().CUID {
		_, err := its.executeLocalBase(op)
		if err != nil { // TODO: if an operation fails to be executed, opID should be rollbacked.
			return err
		}
	} else {
		its.executeRemoteBase(op)
	}
	return nil
}

// SetNextOpID proceeds the operation ID next.
func (its *BaseDatatype) SetNextOpID(op iface.Operation) {
	op.SetID(its.opID.Next())
}

// GetType returns the type of this datatype.
func (its *BaseDatatype) GetType() model.TypeOfDatatype {
	return its.TypeOf
}

// GetState returns the state of this datatype.
func (its *BaseDatatype) GetState() model.StateOfDatatype {
	return its.state
}

// SetOpID sets the operation ID.
func (its *BaseDatatype) SetOpID(opID *model.OperationID) {
	its.opID = opID
}

// GetKey returns the key.
func (its *BaseDatatype) GetKey() string {
	return its.Key
}

// GetDUID returns DUID.
func (its *BaseDatatype) GetDUID() string {
	return its.id
}

// SetDUID sets the DUID.
func (its *BaseDatatype) SetDUID(duid string) {
	its.id = duid
}

// SetState sets the state of this datatype.
func (its *BaseDatatype) SetState(state model.StateOfDatatype) {
	its.state = state
}

func (its *BaseDatatype) GetOpID() *model.OperationID {
	return its.opID
}

// GetMeta returns the binary of metadata of the datatype.
func (its *BaseDatatype) GetMeta() ([]byte, errors.OrdaError) {
	meta := model.DatatypeMeta{
		Key:    its.Key,
		DUID:   its.id,
		OpID:   its.opID,
		TypeOf: its.TypeOf,
	}
	metab, err := json.Marshal(&meta)
	if err != nil {
		return nil, errors.DatatypeMarshal.New(its.ctx.L(), meta)
	}
	return metab, nil
}

// SetMeta sets the metadata with binary metadata.
func (its *BaseDatatype) SetMeta(meta []byte) errors.OrdaError {
	m := model.DatatypeMeta{}
	if err := json.Unmarshal(meta, &m); err != nil {
		return errors.DatatypeMarshal.New(its.ctx.L(), string(meta))
	}
	its.Key = m.Key
	its.id = m.DUID
	its.opID = m.OpID
	its.TypeOf = m.TypeOf
	return nil
}

func (its *BaseDatatype) L() *log.OrdaLog {
	return its.ctx.L()
}

func (its *BaseDatatype) GetSummary() string {
	return fmt.Sprintf("%s(%s)", utils.MakeDefaultShort(its.Key), its.id)
}
