import { RealmType } from "./common";
import { Realm } from "./common/Realm";
import { MagnatesRealm } from "./magnates/MagnatesRealm";

export class SimCompanies {
  static loadRealm(realm: RealmType): Realm {
    return realm === RealmType.Magnates
      ? new MagnatesRealm("http://localhost:8080")
      : null!;
  }
}
