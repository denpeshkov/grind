use std::iter::Peekable;

pub struct Merged<L, R>
where
    L: Iterator,
    R: Iterator,
{
    left: Peekable<L>,
    right: Peekable<R>,
}

impl<L, R> Merged<L, R>
where
    L: Iterator,
    R: Iterator,
{
    pub fn new(l: L, r: R) -> Self {
        Merged {
            left: l.peekable(),
            right: r.peekable(),
        }
    }
}

impl<L, R> Iterator for Merged<L, R>
where
    L: Iterator,
    R: Iterator<Item = L::Item>,
    L::Item: Ord,
{
    type Item = L::Item;

    fn next(&mut self) -> Option<Self::Item> {
        match (self.left.peek(), self.right.peek()) {
            (None, None) => None,
            (None, Some(_)) => self.right.next(),
            (Some(_), None) => self.left.next(),
            (Some(l), Some(r)) => {
                if l <= r {
                    self.left.next()
                } else {
                    self.right.next()
                }
            }
        }
    }
}

#[cfg(test)]
mod tests {
    use super::*;
    use quickcheck_macros::quickcheck;

    #[quickcheck]
    fn proptest_merged(mut l: Vec<i32>, mut r: Vec<i32>) -> bool {
        l.sort();
        r.sort();
        Merged::new(l.iter(), r.iter()).is_sorted()
    }
}
