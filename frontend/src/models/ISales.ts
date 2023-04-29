import {AdminInterface} from "./IAdmin";
import {SellInterface} from "./ISell";
import {TradeInterface} from "./ITrade";

export interface SalesInterface {
	IDsales: 		number,
    DayTotal:   number,
    MonthTotal: number,

    AdminID:    number,
    Admin:      AdminInterface,

    SellID:     number,
    Sell:       SellInterface,

    TradeID:     number,
    Trade:       TradeInterface,
}