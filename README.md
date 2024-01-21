The objective of this study was to assess if Go can independently serve as a comprehensive programming language for both backend development for the engineering team and for statistical tasks conducted by the data science team. To evaluate this, the Anscombe quartet dataset was utilized to perform statistical testing across three programming languages, Python, R and Go. Benchmark studies were conducted with these languages to establish the average computation speed, using 1000 iterations of each test.

Python and R have well established footings in the statistical computation space. Runing the code on the Anscombe dataset with R yields very detailed statistical reports and requires only a few lines of code. Included in the reports are metrics such as correlation coefficients, r-squared values, p-value, as well as standard error and residual analysis among others. Benchmarking the code yielded 0.1075 seconds on average to run. Python demonstrated a similar scripting experience with few lines required for output but generated an even more detailed report giving valuable information such as kurtosis and skew on top of the metrics seen with R. Runing a benchmark with Python gave an average of 0.0255 seconds to run.
Taking the previous data into consideration it becomes clear that Go, in comparison to Python and R, does not facilitate the production of statistical results as swiftly or intuitively. Management recommended the montanflynn/stats package for the analysis of the Anscombe Quartet. This package is equipped with some useful functions that allow for basic statistical analysis, such as standard deviation, mean, median and mode, and even probability density functions as well as many others. However, the user experience when delving into linear regression was not favorable to a new user. Indeed this stats package does have LinearRegression() as an option, however, the upfront code necessary to run this function was more extensive than seen with Python or R. Also, output obtained using the Anscombe Quartet appears to be a slice of stat.Coordinates that give predictions based on the linear model fitted to the data. This discovery highlighted a lack of clarity regarding how to extract vital statistical measures, such as correlation coefficients, the line of best fit or r-squared, using the specified Go statistics package. It appears that this package does not offer comprehensive statistical reporting capabilities that are necessary for proper data analysis. To ensure that the data was being read into the program correctly a manual linear regression calculation was performed in tandem to compare with the expected results of the Anscombe data. The manual linear regression calculation did indeed give the correct coefficients seen in the regression line. In addition, R-squared values are not directly supported by the montanaflynn/stats package. The R-squared value was obtained by squaring the Correlation() function from montanaflynn/stats, but this approach is not feasible with more complex multiple linear regression calculations that may arise. Therefore, it is recommended that if management wants to go forward with the Go programming language across the company, a different stats package should be explored. Benchmarking the stats calculations yielded a 0.146 second result, which shows a slower average among the three programming languages.

In conclusion, Python and R offer simplistic and streamlined statistical analysis frameworks that can produce detailed reports on the dataset of interest. For that reason, it is not my recommendation to move forward with Go exclusively within the company. Go can be used for statistical programming, however another package may be worth considering given the amount of time taken reworking Go code to try to achieve the same result seen with Python and R on the Anscombe dataset. If time is not a constraint, getting the data scientists up to speed using statistics with Go is feasible but code necessary to produce the desired result will increse with the current stats package of interest.