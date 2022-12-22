import { PrefixsInterface } from "./IPrefix";
import { GendersInterface } from "./IGender";
import { PolicingsInterface } from "./IPolicing";

export interface OfficerInterface {
    ID: number,
    Name: string,
    Date_of_birth: string,
    Age: number,
    Phone: string,
    Address: string,
    ID_card: number,
    
    Prefix_ID: number,
    Prefix: PrefixsInterface,

    Gender_ID: number,
    Gender: GendersInterface,

    Policing_ID: number,
    Policing: PolicingsInterface,   
  }