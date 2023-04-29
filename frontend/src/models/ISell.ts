import {AdminInterface} from "./IAdmin";
import {UserInterface} from "./IUser";
import {UnitInterface} from "./IUnit";
import {TradeInterface} from "./ITrade";
import {DrugInterface} from "./IDrug";

export interface SellInterface {
	ID: 		number,
    Quantity:   number,
    Cost: 		number,
	Total: 	    number,
	Payment:  	number,
    SellTime:   Date | null,


    AdminID:    number,
    Admin:      AdminInterface,

    UserID:     number,
    User:       UserInterface,

    UnitID:     number,
    Unit:       UnitInterface,

    TradeID:    number,
    Trade:      TradeInterface,

    DrugID:     number,
    Drug:       DrugInterface,
}