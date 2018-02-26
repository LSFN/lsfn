pub mod lsfn_api {
    include!(concat!(env!("OUT_DIR"), "/lsfn.protocol.rs"));
}

use uuid::Uuid;
use lsfn_api::ShipDescription;


struct Ship {
    id: Uuid,
    name: String,
    physics_body: RigidBody,
    
}

pub trait DescribableShip {
    fn describe(&self) -> ShipDescription;
}

impl DescribableShip for Ship {
    fn describe(&self) -> ShipDescription {

    }
}
