use std::time::SystemTime;
use std::time::UNIX_EPOCH;

use clap::Parser;
use clap::ValueEnum;

#[derive(Parser)]
#[command(version)]
struct Args {
    #[arg(long, short, value_enum)]
    precision: Option<Precision>,
}

#[derive(Clone, Copy, ValueEnum)]
enum Precision {
    Millis,
    Secs,
}

fn main() {
    let args = Args::parse();
    let epoch = SystemTime::now().duration_since(UNIX_EPOCH).unwrap();
    match args.precision.unwrap_or(Precision::Secs) {
        Precision::Millis => println!("{}", epoch.as_millis()),
        Precision::Secs => println!("{}", epoch.as_secs()),
    }
}
