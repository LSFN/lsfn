extern crate tokio;
extern crate tokio_io;
extern crate futures;
extern crate bytes;
extern crate prost;
#[macro_use]
extern crate prost_derive;
extern crate uuid;

use tokio::executor::current_thread;
use tokio::net::TcpListener;
use tokio_io::io;
use futures::{Future, Stream};
use prost::Message;
use uuid::Uuid;

pub mod lsfn_api {
    include!(concat!(env!("OUT_DIR"), "/lsfn.protocol.rs"));
}

use lsfn_api::{Welcome, ShipDescription};

fn main() {
    let addr = "127.0.0.1:6142".parse().unwrap();
    let listener = TcpListener::bind(&addr).unwrap();

    let server = listener.incoming().for_each(|socket| {
        println!("accepted socket; addr={:?}", socket.peer_addr().unwrap());

        let ship = Welcome { ship: Some(ShipDescription {
            id: Uuid::new_v4().to_string(),
            name: "HMS Edge of Reason".to_string(),
            controls: vec![],
            sensors: vec![],
        }) };

        let mut ship_message = vec![];
        ship.encode(&mut ship_message)?;
        let connection = io::write_all(socket, ship_message)
            .then(|res| {
                println!("wrote message; success={:?}", res.is_ok());
                Ok(())
            });

        // Spawn a new task that processes the socket:
        current_thread::spawn(connection);

        Ok(())
    })
        .map_err(|err| {
            // All tasks must have an `Error` type of `()`. This forces error
            // handling and helps avoid silencing failures.
            //
            // In our example, we are only going to log the error to STDOUT.
            println!("accept error = {:?}", err);
        });

    current_thread::run(|_| {
        // Now, the server task is spawned.
        current_thread::spawn(server);

        println!("server running on localhost:6142");
    });
}
