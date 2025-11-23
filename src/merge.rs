use std::{
    fmt::{self, Debug},
    iter::Peekable,
};

/// An iterator that merges two sorted iterators into one sorted iterator.
pub struct Merge<L, R>
where
    L: Iterator,
    R: Iterator,
{
    left: Peekable<L>,
    right: Peekable<R>,
}

impl<L, R> Clone for Merge<L, R>
where
    L: Iterator + Clone,
    L::Item: Clone,
    R: Iterator + Clone,
    R::Item: Clone,
{
    fn clone(&self) -> Self {
        Self {
            left: self.left.clone(),
            right: self.right.clone(),
        }
    }
}

impl<L, R> Debug for Merge<L, R>
where
    L: Iterator + Debug,
    L::Item: Debug,
    R: Iterator + Debug,
    R::Item: Debug,
{
    fn fmt(&self, f: &mut fmt::Formatter<'_>) -> fmt::Result {
        f.debug_struct("Merge")
            .field("left", &self.left)
            .field("right", &self.right)
            .finish()
    }
}

impl<L, R> Merge<L, R>
where
    L: Iterator,
    R: Iterator,
{
    /// Creates an iterator that merges two sorted iterators into one sorted iterator.
    pub fn new(left: L, right: R) -> Self {
        Self {
            left: left.peekable(),
            right: right.peekable(),
        }
    }
}

impl<L, R> Iterator for Merge<L, R>
where
    L: Iterator,
    L::Item: Ord,
    R: Iterator<Item = L::Item>,
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
        Merge::new(l.iter(), r.iter()).is_sorted()
    }
}
