extern crate nalgebra as na;

use na::Point3;
use uuid::Uuid;

struct Thruster {
    id: Uuid

}

impl Sensor<Point3> for GalacticPositionSensor {
    fn id(&self) -> Uuid {
        return self.id;
    }
    fn name(&self) -> string {
        return "Galactic Position Sensor"
    }
    fn read_sensor_data(&self, ship: Ship) -> Point3 {
        return ship.physics_body.position()
    }
}
