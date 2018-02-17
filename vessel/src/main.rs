extern crate glium;
extern crate glium_text;
extern crate cgmath;
extern crate bytes;
extern crate prost;
#[macro_use]
extern crate prost_derive;

use std::thread;
use std::time::Duration;
use bytes::Buf;
use std::io::Cursor;
use glium::Surface;
use glium::glutin;
use prost::Message;
use lsfn_api::{Welcome, ShipDescription, ControlDescription, SensorDescription};
use lsfn_api::control_description::ControlType;
use lsfn_api::sensor_description::SensorType;

pub mod lsfn_api {
    include!(concat!(env!("OUT_DIR"), "/lsfn.protocol.rs"));
}

use std::io::prelude::*;
use std::net::TcpStream;


fn main() {

    let mut stream = TcpStream::connect("127.0.0.1:6142").unwrap();

    let mut buf = [0; 512];
    let num = stream.read(&mut buf).unwrap();

    println!("Read {:?} bytes", num);
    println!("Raw: {:?}", buf.to_vec());

    let w = Welcome::decode(&buf[0..num].to_vec()).expect("Failed to reach server");

    println!("Response: {:?}", w);

    use glium::DisplayBuild;

    let display = glutin::WindowBuilder::new().with_dimensions(1024, 768).build_glium().unwrap();
    let system = glium_text::TextSystem::new(&display);

    let font = glium_text::FontTexture::new(&display, &include_bytes!("font/mainframe/mainframe-opto.ttf")[..], 70).unwrap();

    let text = glium_text::TextDisplay::new(&system, &font, &w.ship.unwrap().name);
    let text_width = text.get_width();
    println!("Text width: {:?}", text_width);

    let sleep_duration = Duration::from_millis(17);

    'main: loop {
        let (w, h) = display.get_framebuffer_dimensions();

        let matrix:[[f32; 4]; 4] = cgmath::Matrix4::new(
            2.0 / text_width, 0.0, 0.0, 0.0,
            0.0, 2.0 * (w as f32) / (h as f32) / text_width, 0.0, 0.0,
            0.0, 0.0, 1.0, 0.0,
            -1.0, -1.0, 0.0, 1.0f32,
        ).into();

        let mut target = display.draw();
        target.clear_color(0.0, 0.0, 0.0, 1.0);
        glium_text::draw(&text, &system, &mut target, matrix, (1.0, 1.0, 0.0, 1.0));
        target.finish().unwrap();

        thread::sleep(sleep_duration);

        for event in display.poll_events() {
            match event {
                glutin::Event::Closed => break 'main,
                _ => ()
            }
        }
    }
}
