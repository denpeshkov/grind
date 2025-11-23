// Sorts slice in ascending order using bubble sort algorithm.
pub fn bubble_sort<T: Ord>(s: &mut [T]) {
    for i in 0..s.len() {
        for j in 1..s.len() - i {
            if s[j - 1] > s[j] {
                s.swap(j - 1, j);
            }
        }
    }
}

// Sorts slice in ascending order using insertion sort algorithm.
pub fn insertion_sort<T: Ord>(s: &mut [T]) {
    for i in 1..s.len() {
        let mut j = i;
        while j > 0 && s[j - 1] > s[j] {
            s.swap(j - 1, j);
            j -= 1
        }
    }
}

// Sorts slice in ascending order using merge sort algorithm.
pub fn merge_sort<T: Ord + Clone>(s: &mut [T]) {
    fn merge<T: Ord + Clone>(s1: &mut [T], s2: &mut [T], aux: &mut Vec<T>) {
        use super::merge::Merge;
        aux.extend(Merge::new(s1.iter().cloned(), s2.iter().cloned()));
    }
    fn sort<T: Ord + Clone>(s: &mut [T], aux: &mut Vec<T>) {
        if s.len() <= 1 {
            return;
        }
        let (l, r) = s.split_at_mut(s.len() / 2);
        sort(l, aux);
        sort(r, aux);
        merge(l, r, aux);
        s.clone_from_slice(aux);
        aux.clear();
    }

    let mut aux = Vec::with_capacity(s.len());
    sort(s, &mut aux);
}

#[cfg(test)]
mod tests {
    use super::*;
    use quickcheck_macros::quickcheck;

    #[quickcheck]
    fn proptest_bubble_sort(mut xs: Vec<i32>) -> bool {
        bubble_sort(&mut xs);
        xs.is_sorted()
    }

    #[quickcheck]
    fn proptest_insertion_sort(mut xs: Vec<i32>) -> bool {
        insertion_sort(&mut xs);
        xs.is_sorted()
    }

    #[quickcheck]
    fn proptest_merge_sort(mut xs: Vec<i32>) -> bool {
        merge_sort(&mut xs);
        xs.is_sorted()
    }
}
