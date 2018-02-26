extern crate nalgebra as na;
extern crate ncollide;
extern crate nphysics3d;

use na::{Point3, Vector3, Translation3};
use ncollide::shape::{Plane, Cuboid, Cone, Cylinder, Ball};
use nphysics3d::world::World;
use nphysics3d::object::RigidBody;
use nphysics_testbed3d::Testbed;

fn make_ship_body() -> RigidBody {
    let geom = Cuboid::new(Vector3::new(10, 10, 10));
    return RigidBody::new_dynamic(geom, 1.0, 0.3, 0.5);
}
