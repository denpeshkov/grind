/// Searches the slice for a given element using linear searching algorithm.
/// Returns an index of the target if it exists.
pub fn linear_search<T: Eq>(s: &[T], target: &T) -> Option<usize> {
    for (i, item) in s.iter().enumerate() {
        if item == target {
            return Some(i);
        }
    }
    None
}

/// Returns the index of the partition point according to the given predicate.
/// The slice is assumed to be partitioned according to the given predicate.
pub fn partition_point<P>(n: usize, mut pred: P) -> usize
where
    P: FnMut(usize) -> bool,
{
    let (mut l, mut r) = (0, n);
    while l < r {
        let m = (l + r) >> 1;
        if pred(m) {
            l = m + 1;
        } else {
            r = m;
        }
    }
    l
}

/// Searches the slice for a given element using binary search algorithm.
/// If the slice is not sorted, the returned result is unspecified.
///
/// If the value is found then [`Result::Ok`] is returned, containing the index of the matching element.
/// If there are multiple matches, then any one of the matches could be returned.
///
/// If the value is not found then [`Result::Err`] is returned, containing
/// the index where a matching element could be inserted while maintaining sorted order.
pub fn binary_search<T: Ord>(s: &[T], target: &T) -> Result<usize, usize> {
    let idx = partition_point(s.len(), |i| s[i] < *target);
    if idx < s.len() && s[idx] == *target {
        return Ok(idx);
    }
    Err(idx)
}

#[cfg(test)]
mod test {
    use super::*;
    use quickcheck_macros::quickcheck;

    #[quickcheck]
    fn proptest_linear_search(xs: Vec<i32>, target: i32) -> bool {
        let idx = linear_search(&xs, &target);
        match idx {
            Some(i) => i < xs.len() && xs[i] == target && xs[..i].iter().all(|x| x != &target),
            None => !xs.contains(&target),
        }
    }

    #[test]
    fn test_partition_point() {
        let b: [i32; 0] = [];
        assert_eq!(partition_point(b.len(), |i| b[i] < 5), 0);

        let b = [4];
        assert_eq!(partition_point(b.len(), |i| b[i] < 3), 0);
        assert_eq!(partition_point(b.len(), |i| b[i] < 4), 0);
        assert_eq!(partition_point(b.len(), |i| b[i] < 5), 1);

        let b = [1, 2, 4, 6, 8, 9];
        assert_eq!(partition_point(b.len(), |i| b[i] < 5), 3);
        assert_eq!(partition_point(b.len(), |i| b[i] < 6), 3);
        assert_eq!(partition_point(b.len(), |i| b[i] < 7), 4);
        assert_eq!(partition_point(b.len(), |i| b[i] < 8), 4);

        let b = [1, 2, 4, 5, 6, 8];
        assert_eq!(partition_point(b.len(), |i| b[i] < 9), 6);

        let b = [1, 2, 4, 6, 7, 8, 9];
        assert_eq!(partition_point(b.len(), |i| b[i] < 6), 3);
        assert_eq!(partition_point(b.len(), |i| b[i] < 5), 3);
        assert_eq!(partition_point(b.len(), |i| b[i] < 8), 5);

        let b = [1, 2, 4, 5, 6, 8, 9];
        assert_eq!(partition_point(b.len(), |i| b[i] < 7), 5);
        assert_eq!(partition_point(b.len(), |i| b[i] < 0), 0);

        let b = [1, 3, 3, 3, 7];
        assert_eq!(partition_point(b.len(), |i| b[i] < 0), 0);
        assert_eq!(partition_point(b.len(), |i| b[i] < 1), 0);
        assert_eq!(partition_point(b.len(), |i| b[i] < 2), 1);
        assert_eq!(partition_point(b.len(), |i| b[i] < 3), 1);
        assert_eq!(partition_point(b.len(), |i| b[i] < 4), 4);
        assert_eq!(partition_point(b.len(), |i| b[i] < 5), 4);
        assert_eq!(partition_point(b.len(), |i| b[i] < 6), 4);
        assert_eq!(partition_point(b.len(), |i| b[i] < 7), 4);
        assert_eq!(partition_point(b.len(), |i| b[i] < 8), 5);

        let data = [-10, -5, 0, 1, 2, 3, 5, 7, 11, 100, 100, 100, 1000, 10000];

        assert_eq!(partition_point(1, |i| i < 1), 1);
        assert_eq!(partition_point(1, |_| false), 0);
        assert_eq!(partition_point(1, |_| true), 1);
        assert_eq!(partition_point(1e9 as usize, |i| i < 991), 991);
        assert_eq!(partition_point(1e9 as usize, |_| false), 0);
        assert_eq!(partition_point(1e9 as usize, |_| true), 1e9 as usize);
        assert_eq!(partition_point(data.len(), |i| data[i] < -20), 0);
        assert_eq!(partition_point(data.len(), |i| data[i] < -10), 0);
        assert_eq!(partition_point(data.len(), |i| data[i] < -9), 1);
        assert_eq!(partition_point(data.len(), |i| data[i] < -6), 1);
        assert_eq!(partition_point(data.len(), |i| data[i] < -5), 1);
        assert_eq!(partition_point(data.len(), |i| data[i] < 3), 5);
        assert_eq!(partition_point(data.len(), |i| data[i] < 11), 8);
        assert_eq!(partition_point(data.len(), |i| data[i] < 99), 9);
        assert_eq!(partition_point(data.len(), |i| data[i] < 100), 9);
        assert_eq!(partition_point(data.len(), |i| data[i] < 101), 12);
        assert_eq!(partition_point(data.len(), |i| data[i] < 10000), 13);
        assert_eq!(partition_point(data.len(), |i| data[i] < 10001), 14);
        assert_eq!(
            partition_point(7, |i| [99, 99, 59, 42, 7, 0, -1, -1][i] > 7),
            4
        );
        assert_eq!(
            partition_point(1e9 as usize, |i| 1e9 as usize - i > 7),
            1e9 as usize - 7
        );
        assert_eq!(partition_point(2e9 as usize, |_| true), 2e9 as usize);
    }

    #[test]
    fn test_binary_search() {
        let b: [i32; 0] = [];
        assert_eq!(binary_search(&b, &5), Err(0));

        let b = [4];
        assert_eq!(binary_search(&b, &3), Err(0));
        assert_eq!(binary_search(&b, &4), Ok(0));
        assert_eq!(binary_search(&b, &5), Err(1));

        let b = [1, 2, 4, 6, 8, 9];
        assert_eq!(binary_search(&b, &5), Err(3));
        assert_eq!(binary_search(&b, &6), Ok(3));
        assert_eq!(binary_search(&b, &7), Err(4));
        assert_eq!(binary_search(&b, &8), Ok(4));

        let b = [1, 2, 4, 5, 6, 8];
        assert_eq!(binary_search(&b, &9), Err(6));

        let b = [1, 2, 4, 6, 7, 8, 9];
        assert_eq!(binary_search(&b, &6), Ok(3));
        assert_eq!(binary_search(&b, &5), Err(3));
        assert_eq!(binary_search(&b, &8), Ok(5));

        let b = [1, 2, 4, 5, 6, 8, 9];
        assert_eq!(binary_search(&b, &7), Err(5));
        assert_eq!(binary_search(&b, &0), Err(0));

        let b = [1, 3, 3, 3, 7];
        assert_eq!(binary_search(&b, &0), Err(0));
        assert_eq!(binary_search(&b, &1), Ok(0));
        assert_eq!(binary_search(&b, &2), Err(1));
        assert!(matches!(binary_search(&b, &3), Ok(1..=3)));
        assert!(matches!(binary_search(&b, &3), Ok(1..=3)));
        assert_eq!(binary_search(&b, &4), Err(4));
        assert_eq!(binary_search(&b, &5), Err(4));
        assert_eq!(binary_search(&b, &6), Err(4));
        assert_eq!(binary_search(&b, &7), Ok(4));
        assert_eq!(binary_search(&b, &8), Err(5));

        let b = [(); usize::MAX];
        assert_eq!(binary_search(&b, &()), Ok(0));
    }
}
