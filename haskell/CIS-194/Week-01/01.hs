toDigits :: Integer -> [Integer]
toDigits n = go n []
  where
    go n xs
      | n <= 0 = xs
      | otherwise = go (n `div` 10) (n `mod` 10 : xs)

toDigitsRev :: Integer -> [Integer]
toDigitsRev n = reverse (toDigits n)
