use uuid::Uuid;

pub trait Control<T> {
    fn id(&self) -> Uuid;
    fn name(&self) -> string;
    fn set_control(&self, ship: Ship, value: T) -> T;
}
