extern crate prost_build;

fn main() {
    prost_build::compile_protos(&["../protocol/env-ship.proto"],
                                &["../"]).unwrap();
}
