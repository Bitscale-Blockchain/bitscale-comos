import { GeneratedType } from "@cosmjs/proto-signing";
import { QueryAllDenomRequest } from "./types/bitscale/bitscale/query";
import { MsgUpdateDenom } from "./types/bitscale/bitscale/tx";
import { BitscalePacketData } from "./types/bitscale/bitscale/packet";
import { MsgUpdateParamsResponse } from "./types/bitscale/bitscale/tx";
import { MsgCreateDenomResponse } from "./types/bitscale/bitscale/tx";
import { Denom } from "./types/bitscale/bitscale/denom";
import { MsgCreateDenom } from "./types/bitscale/bitscale/tx";
import { MsgDeleteDenom } from "./types/bitscale/bitscale/tx";
import { NoData } from "./types/bitscale/bitscale/packet";
import { MsgDeleteDenomResponse } from "./types/bitscale/bitscale/tx";
import { QueryParamsResponse } from "./types/bitscale/bitscale/query";
import { MsgUpdateDenomResponse } from "./types/bitscale/bitscale/tx";
import { QueryGetDenomResponse } from "./types/bitscale/bitscale/query";
import { QueryAllDenomResponse } from "./types/bitscale/bitscale/query";
import { GenesisState } from "./types/bitscale/bitscale/genesis";
import { MsgUpdateParams } from "./types/bitscale/bitscale/tx";
import { Params } from "./types/bitscale/bitscale/params";
import { QueryParamsRequest } from "./types/bitscale/bitscale/query";
import { QueryGetDenomRequest } from "./types/bitscale/bitscale/query";

const msgTypes: Array<[string, GeneratedType]>  = [
    ["/bitscale.bitscale.QueryAllDenomRequest", QueryAllDenomRequest],
    ["/bitscale.bitscale.MsgUpdateDenom", MsgUpdateDenom],
    ["/bitscale.bitscale.BitscalePacketData", BitscalePacketData],
    ["/bitscale.bitscale.MsgUpdateParamsResponse", MsgUpdateParamsResponse],
    ["/bitscale.bitscale.MsgCreateDenomResponse", MsgCreateDenomResponse],
    ["/bitscale.bitscale.Denom", Denom],
    ["/bitscale.bitscale.MsgCreateDenom", MsgCreateDenom],
    ["/bitscale.bitscale.MsgDeleteDenom", MsgDeleteDenom],
    ["/bitscale.bitscale.NoData", NoData],
    ["/bitscale.bitscale.MsgDeleteDenomResponse", MsgDeleteDenomResponse],
    ["/bitscale.bitscale.QueryParamsResponse", QueryParamsResponse],
    ["/bitscale.bitscale.MsgUpdateDenomResponse", MsgUpdateDenomResponse],
    ["/bitscale.bitscale.QueryGetDenomResponse", QueryGetDenomResponse],
    ["/bitscale.bitscale.QueryAllDenomResponse", QueryAllDenomResponse],
    ["/bitscale.bitscale.GenesisState", GenesisState],
    ["/bitscale.bitscale.MsgUpdateParams", MsgUpdateParams],
    ["/bitscale.bitscale.Params", Params],
    ["/bitscale.bitscale.QueryParamsRequest", QueryParamsRequest],
    ["/bitscale.bitscale.QueryGetDenomRequest", QueryGetDenomRequest],
    
];

export { msgTypes }