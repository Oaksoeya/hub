package types

import (
	"github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrorMarshal              = errors.Register(ModuleName, 101, "error occurred while marshalling")
	ErrorUnmarshal            = errors.Register(ModuleName, 102, "error occurred while unmarshalling")
	ErrorUnknownMsgType       = errors.Register(ModuleName, 103, "unknown message type")
	ErrorUnknownQueryType     = errors.Register(ModuleName, 104, "unknown query type")
	ErrorInvalidField         = errors.Register(ModuleName, 105, "invalid field")
	ErrorDuplicateProvider    = errors.Register(ModuleName, 106, "duplicate provider")
	ErrorProviderDoesNotExist = errors.Register(ModuleName, 107, "provider does not exist")
)
