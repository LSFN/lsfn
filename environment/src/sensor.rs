use uuid::Uuid;

pub trait Sensor<T> {
    fn id(&self) -> Uuid;
    fn name(&self) -> string;
    fn read_sensor_data(&self, ship: Ship) -> T;
}
