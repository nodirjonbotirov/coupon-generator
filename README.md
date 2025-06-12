# Coupon Generator

Coupon Generator is a Go application for generating unique coupon codes based on customizable patterns. It supports exporting generated coupons to a CSV file for easy integration with marketing, sales, or promotional campaigns.

## Features
- **Customizable Patterns:** Define coupon code formats using placeholders:
  - `A` for uppercase letters (excluding ambiguous ones)
  - `D` for digits (excluding ambiguous ones)
  - `S` for symbols
  - Any other character is used as a literal (e.g., dashes or colons)
- **Unique Code Generation:** Ensures all generated coupons are unique.
- **Export to CSV:** Save generated coupons to a CSV file for further use.
- **Configurable Quantity:** Specify how many unique coupons to generate.

## Usage
1. **Configure the Pattern and Quantity:**
   - Edit the `pattern` and `couponCount` variables in `main.go` to set your desired coupon format and number of codes.
2. **Run the Application:**
   ```sh
   go run main.go
   ```
3. **View/Export Coupons:**
   - Generated coupons are printed to the console and saved to `coupons.csv` in the project directory.

## Example
A pattern of `AADD` generates codes like `JK12`, `GH34`, etc., where `A` is a letter and `D` is a digit.

## Pattern Placeholders
- `A`: Uppercase letter (A, B, C, D, G, H, J, K, L, M, N, P, Q, R, S, T, W, X, Y, Z)
- `D`: Digit (0, 1, 2, 3, 4, 5, 7, 8)
- `S`: Symbol (!, @, #, $, %)

## Output
- **Console:** All generated coupons are printed.
- **CSV File:** Coupons are saved in `coupons.csv`, one per line.

## Requirements
- Go 1.18 or later

## License
MIT License
