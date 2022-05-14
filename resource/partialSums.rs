fn main() {
	let input: Vec<u16> = std::io::stdin().lines().map(|line| line.unwrap().parse::<u16>().unwrap()).collect();

	let mut sums = Vec::<u16>::with_capacity(input.len());
	let mut running_total = 0u16;

	for value in input {
		running_total += value;
		sums.push(running_total);
	}

	for value in sums.iter().flat_map(|value| value.to_le_bytes()) {
		print!("\\x{:02x}", value)
	}
}