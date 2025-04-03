use std::io::stdin;
use std::rc::Rc;

use chrono::NaiveTime;

type Error = String;

fn main() {
    if let Err(e) = execute() {
        println!("{e}");
    };
}

fn execute() -> Result<(), Error> {
    let lines = stdin()
        .lines()
        .collect::<Result<Vec<_>, _>>()
        .map_err(|_| "failed to read stdin".to_string())?
        .into_iter()
        .filter(|s| !s.chars().all(|c| c.is_ascii_whitespace()))
        .collect::<Vec<_>>();

    let days: i32 = lines.len().try_into().unwrap();

    let time_pairs = lines
        .iter()
        .map(|line| {
            line.split(' ')
                .map(|s| NaiveTime::parse_from_str(s, "%H:%M"))
                .collect::<Result<Rc<_>, _>>()
        })
        .collect::<Result<Rc<_>, _>>()
        .map_err(|_| "time format not correct".to_string())?
        .iter()
        .map(|times| -> Result<(NaiveTime, NaiveTime), Error> {
            let start_time = times.first().ok_or("time format not correct".to_string())?;
            let end_time = times.get(1).ok_or("time format not correct".to_string())?;
            Ok((*start_time, *end_time))
        })
        .collect::<Result<Vec<_>, _>>()?;
    let work_time: i32 = time_pairs
        .into_iter()
        .map(|(start_time, end_time)| end_time - start_time)
        .map(|delta| delta.num_minutes())
        .into_iter()
        .sum::<i64>()
        .try_into()
        .unwrap();

    let average = work_time as f64 / days as f64 / 60f64;
    let expected = days * 12 * 60;
    let diff = work_time - expected;

    println!("days: {days}");
    println!("average: {average:.4} hrs");
    println!("diff: {diff:+} mins");
    Ok(())
}
