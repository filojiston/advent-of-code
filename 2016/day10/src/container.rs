#[derive(Debug)]
pub struct Container {
    pub id: String,
    low: Option<u32>,
    high: Option<u32>,
}

impl Container {
    pub fn new(id: String) -> Self {
        Container {
            id: id,
            low: None,
            high: None,
        }
    }

    pub fn give(&mut self, value: u32) {
        if self.low.is_none() && self.high.is_none() {
            self.low = Some(value);
        } else if self.low.is_some() && self.high.is_none() {
            if value < self.low() {
                self.high = Some(self.low());
                self.low = Some(value);
            } else {
                self.high = Some(value);
            }
        } else if self.low.is_some() && self.high.is_some() {
            if value < self.low() {
                self.high = Some(self.low());
                self.low = Some(value);
            } else if value > self.high() {
                self.high = Some(value);
            } else {
                self.low = Some(value);
            }
        }
    }

    pub fn reset(&mut self) {
        self.low = None;
        self.high = None;
    }

    pub fn is_ready(&self) -> bool {
        self.low.is_some() && self.high.is_some()
    }

    pub fn low(&self) -> u32 {
        self.low.unwrap()
    }

    pub fn high(&self) -> u32 {
        self.high.unwrap()
    }
}
