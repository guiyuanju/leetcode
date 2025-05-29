fn main() {
    let mut s = "axc";
    let mut t = "ahbgdc";
    println!("{}", is_subsequence(s.to_string(), t.to_string()));
}

pub fn is_subsequence(s: String, t: String) -> bool {
    let s = s.as_bytes();
    let t = t.as_bytes();
    let mut i = 0;
    let mut j = 0;
    while i < s.len() && j < t.len() {
        if s[i] == t[j] {
            i += 1;
        }
        j += 1;
    }
    i >= s.len()
}
