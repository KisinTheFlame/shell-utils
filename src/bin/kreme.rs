use std::fs;
use std::path::PathBuf;

use clap::Parser;
use regex::Regex;

type Error = String;

#[derive(Parser)]
struct Args {
    #[arg(long)]
    dry_run: bool,
    #[arg(short, long)]
    dir: PathBuf,
    pattern: String,
    replacement: String,
}

fn main() {
    if let Err(e) = execute() {
        println!("kreme: {e}");
    }
}

fn execute() -> Result<(), Error> {
    let Args {
        dry_run,
        dir,
        pattern,
        replacement,
    } = Args::parse();
    let files = fs::read_dir(dir.clone())
        .map_err(|_| format!("{}: no such directory", dir.to_string_lossy()))?
        .collect::<Result<Vec<_>, _>>()
        .map_err(|_| "failed to walk directory")?;
    let regex = Regex::new(pattern.as_str()).map_err(|e| format!("regex invalid: {e}"))?;
    for file in files {
        let old_name = file.file_name().into_string().unwrap();
        let new_name = regex.replace(&old_name, replacement.as_str()).into_owned();
        if dry_run {
            println!("{old_name} ==> {new_name}");
        } else {
            fs::rename(dir.join(old_name.clone()), dir.join(new_name))
                .map_err(|_| format!("cannot apply on file {old_name}"))?;
        }
    }
    Ok(())
}
