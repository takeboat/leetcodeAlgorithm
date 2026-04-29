struct BrowserHistory {
    history: Vec<String>,
    cur: usize,
}

impl BrowserHistory {
    fn new(homepage: String) -> Self {
        BrowserHistory {
            history: vec![homepage],
            cur: 0,
        }
    }

    fn visit(&mut self, url: String) {
        self.cur += 1;
        self.history.truncate(self.cur);
        self.history.push(url);
    }

    fn back(&mut self, steps: i32) -> String {
        self.cur = (self.cur - steps as usize).max(0);
        self.history[self.cur].clone()
    }

    fn forward(&mut self, steps: i32) -> String {
        self.cur = (self.cur + steps as usize).min(self.history.len() - 1);
        self.history[self.cur].clone()
    }
}
fn main() {}
