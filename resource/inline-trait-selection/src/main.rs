#![feature(io_read_to_string)]
// fn main() {
// 	let input: Vec<u16> = std::io::stdin().lines().map(|line| line.unwrap().parse::<u16>().unwrap()).collect();

// 	let mut sums = Vec::<u16>::with_capacity(input.len());
// 	let mut running_total = 0u16;

// 	for value in input {
// 		running_total += value;
// 		sums.push(running_total);
// 	}

// 	for value in sums.iter().flat_map(|value| value.to_le_bytes()) {
// 		print!("\\x{:02x}", value)
// 	}
// }

fn partial_sums(values: &[u16]) -> Vec<u16> {
    let mut sums = Vec::with_capacity(values.len());
    let mut running_total = 0;

    for value in values {
        running_total += value;
        sums.push(running_total);
    }

    return sums;
}

use std::str::FromStr;

use std::fmt::Display;

#[derive(Debug)]
struct Trait {
    id: u8,
    name: String,
}

fn bisect<T: PartialOrd + Display>(start: usize, end: usize, partial_sums: &[T], values: &[Trait]) {
    if start >= partial_sums.len() {
        println!("require(false, \"Value out of bounds\");");
        return;
    }
    if start > end {
        return;
    }
    if start == end {
        let value = &values[start];
        println!("return ({}, \"{}\");", value.id, value.name);
    } else {
        let boundary = (start + end) / 2;

        println!("if ({} < pick) {{", partial_sums[boundary]);
        bisect(boundary + 1, end, partial_sums, values);
        println!("}} else {{");
        bisect(start, boundary, partial_sums, values);
        println!("}}");
    }
}

fn main() -> anyhow::Result<()> {
    let mut stdin = std::io::stdin().lock();

    let stdin = std::io::read_to_string(&mut stdin)?;
    let value = serde_json::Value::from_str(&stdin)?;
    let array = value
        .as_array()
        .ok_or_else(|| anyhow::anyhow!("Not an array"))?;

    let mut traits = Vec::with_capacity(array.len());
    let mut weights = Vec::with_capacity(array.len());
    for value in array {
        traits.push(Trait {
            id: u8::try_from(value["id"].as_u64().unwrap())?,
            name: value["name"].as_str().unwrap().to_owned(),
        });
        weights.push(u16::try_from(value["weight"].as_u64().unwrap())?);
    }
    let partial_sums = partial_sums(&weights);

    println!("function getTrait(uint32 pick) internal pure returns (uint8 id, bytes28 name) {{");
    bisect(0, traits.len(), &partial_sums, &traits);
    println!("}}");

    Ok(())
}
