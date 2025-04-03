use std::io::BufRead;
use std::io::stdin;
use std::rc::Rc;

use clap::Parser;
use clap::Subcommand;
use regex::Regex;

type Error = String;

#[derive(Parser)]
#[command(version)]
struct Args {
    #[command(subcommand)]
    command: Command,
}

#[derive(Subcommand)]
enum Command {
    Trim,
    Trim2,
    Split { seperator: String },
    Subst { regex: String, replacement: String },
    Rev,
    Head { num: usize },
    Tail { num: usize },
}

fn main() {
    if let Err(e) = execute() {
        println!("{e}");
    }
}

fn execute() -> Result<(), Error> {
    let args = Args::parse();

    let lines = stdin()
        .lock()
        .lines()
        .collect::<Result<Rc<_>, _>>()
        .map_err(|_| "failed to read stdin")?;

    match args.command {
        Command::Trim2 => {
            let output = lines.join("\n").trim().to_string();
            print!("{output}");
        }
        Command::Trim => {
            let output = lines.iter().map(|s| s.trim()).collect::<Rc<_>>().join("\n");
            println!("{output}");
        }
        Command::Split { seperator } => {
            let re = Regex::new(seperator.as_str()).map_err(|e| format!("regex invalid: {e}"))?;
            let output = lines
                .iter()
                .flat_map(|s| re.split(s))
                .collect::<Rc<_>>()
                .join("\n");
            println!("{output}");
        }
        Command::Subst { regex, replacement } => {
            let re = Regex::new(regex.as_str()).map_err(|e| format!("regex invalid: {e}"))?;
            let output = lines
                .iter()
                .map(|s| re.replace_all(s, replacement.as_str()))
                .collect::<Rc<_>>()
                .join("\n");
            println!("{output}");
        }
        Command::Rev => {
            let output = lines
                .iter()
                .map(|s| s.chars().rev().collect::<String>())
                .collect::<Rc<_>>()
                .join("\n");
            println!("{output}");
        }
        Command::Head { num } => {
            let output = lines
                .iter()
                .map(|s| s.chars().take(num).collect::<String>())
                .collect::<Rc<_>>()
                .join("\n");
            println!("{output}");
        }
        Command::Tail { num } => {
            let output = lines
                .iter()
                .map(|s| {
                    s.chars()
                        .rev()
                        .take(num)
                        .collect::<Vec<_>>()
                        .into_iter()
                        .rev()
                        .collect::<String>()
                })
                .collect::<Rc<_>>()
                .join("\n");
            println!("{output}");
        }
    };
    Ok(())
}
