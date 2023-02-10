use lazy_static::lazy_static;
use regex::Regex;

#[derive(Debug, Clone)]
pub struct Instruction {
    pub source: String,
    pub target_low: String,
    pub target_high: String,
    pub value: Option<u32>,
}

lazy_static! {
    static ref BOT_PARSER: Regex = Regex::new(r"(?P<source>bot (\d+)) gives low to (?P<target_low>(bot|output) (\d+)) and high to (?P<target_high>(bot|output) (\d+))").unwrap();
    static ref VALUE_PARSER: Regex = Regex::new(r"value (?P<value>(\d+)) goes to (?P<target>bot (\d+))").unwrap();
}

impl From<&str> for Instruction {
    fn from(line: &str) -> Self {
        let tokens = line.split_whitespace().collect::<Vec<_>>();
        if tokens[0] == "value" {
            let captures = VALUE_PARSER.captures(line).ok_or("no captures").unwrap();
            let value = captures
                .name("value")
                .unwrap()
                .as_str()
                .parse::<u32>()
                .unwrap();
            let bot_id = captures.name("target").unwrap().as_str().to_string();
            Instruction {
                source: bot_id.clone(),
                target_low: bot_id.clone(),
                target_high: bot_id.clone(),
                value: Some(value),
            }
        } else if tokens[0] == "bot" {
            let captures = BOT_PARSER.captures(line).ok_or("no captures").unwrap();
            let source_id = captures.name("source").unwrap().as_str().to_string();
            let target_low_id = captures.name("target_low").unwrap().as_str().to_string();
            let target_high_id = captures.name("target_high").unwrap().as_str().to_string();
            Instruction {
                source: source_id,
                target_low: target_low_id,
                target_high: target_high_id,
                value: None,
            }
        } else {
            panic!("Invalid instruction");
        }
    }
}
