![Banner](../media/gostack_SmallerTransparent.png)

 MM/DD/YYYY
 06/18/2022 — Development begins

 <h1>Current Releases</h1>

 <h2><b>v0.8.0</b> — <i>Alpha Release</i> — 07/24/2022 — Prerelease A</h2>

* ***gostack***'s initial non-working release
* Initialized most function, meaning you are able to call them from other scripts

 <h2><b>v0.8.5</b> — <i>Alpha Release</i> — 08/24/2022 — Prerelease B</h2>

* Significantly improved **gostack** documentation
* Reworked many ***gostack*** functions' functionality

 <h2><b>v0.9.0</b> — <i>Alpha Release</i> — 08/28/2022 — Prerelease C</h2>

* Implemented the **stack.Equals()** and **card.Equals()** functions
* Redesigned significant portions of ***gostack*** memory management

 <h2><b>v1.0.0</b> — <i>Alpha Release</i> — 10/16/2022 — First working implementation</h2>

* ***gostack***'s initial working release
* Every initially-planned function is fully implemented and case-tested

 <h2><b>v1.0.4</b> — <i>Beta Release</i> — 11/13/2022 — Post-release polish</h2>

* Community Discord server, Patreon page, and buymeaoffee is created
* The **Height** property of **Stacks** is renamed to **Height** to improve conceptual clarity
* **MakeSubstack()** function added
* **card.SwitchKeyVal()** function added
* **stack.SwitchKeysVals()** function added
* **stack.Filter()** functions added
* Allow simplified Lambda function parameterizations
* Added more DEFEREFERENCE options for improved flexibility
* Created benchmark tests
* Added LICENSE file
* Added DEREFERENCE_X support to card.Equals and stack.Equals
* **csvToStackMatrix() (\*Stack)** function added
* **stack.ToCSV() ([][]string)** function added
* Abstracted ***stack.Lambda()*** input function
* Added ***stack.Coordinates()*** and ***stack.CoordinatesMany()*** to ***gostack***
* Add coordinates to Lambda functions in otherInfo using forwardpropagation
* Add function samples to each function documentation file
* ***gostack*** functions parameter naming has been adjusted to be more consistent
* ***gostack***'s functions' default parameter values have been adjusted to make more intuitive sense
* ***gostack*** receives a documentation overhaul
* Removed all files that won't be in the full release (e.g., `unaddedcases.txt`)
* ***gostack*** is finally opened and marketed to the public

 <h1>Planned Releases</h1>

 <h2><b>v1.0.5</b> — <i>Beta Release</i> — 01/26/2023 — Patch A</h2>

* stack.Transpose() added
* stack.Shape() => returns type Stack
* stack.StripStackMatrix => stack.DimensionalityReduce
* FIND_Coords added
* Added testing tutorial
* Updated MakeStackMatrix's `matrixShape` parameter to accept Stack arguments
* Fixed a bug with the ToMatrix and ToArray functions not properly handling RETURN_Card and RETURN_Stack arguments

 <h2><b>v1.0.6</b> — <i>Beta Release</i> — ?/?/2023 — Patch B</h2>

* Add FIND_Depth enum
* Implement a height alternative to the depth parameter
* More strictly implement design-by-contract principles in function documentation to facilitate error-catching for invalid arguments
* Further simplify function parameterizations
* Make all lambda function calls have the same structure (i.e., no more simplified version for `replaceWith` func)
* Mathematical functions added (mean, sum, product, add, subtract, multiply, divide, matrix multiply, transpose)
* Optimize to improve performance
  * Add more benchmarks
  * Remove all dependencies on `gogenerics` since this significantly detriments efficiency
* Add GetAddress function to simplify Sprintf %p (possibly just to gogenerics)
* Implement feedback from the ***gostack*** community
    * Bug fixes
    * Function restructuring
    * Documentational improvements
    * Conceptual improvements
* Made `CSVToStackMatrix` less prone to producing errors

 <h2><b>v1.0.8</b> — <i>Beta Release</i> — ?/?/???? — Patch C</h2>

* Implement feedback from the ***gostack*** community
    * Bug fixes
    * Function restructuring
    * Documentational improvements
    * Conceptual improvements

 <h2><b>v1.1.0</b> — <i>Full release</i> — ?/?/???? — Patch D</h2>

* Implement feedback from the ***gostack*** community
    * Bug fixes
    * Function restructuring
    * Documentational improvements
    * Conceptual improvements

 <h2><b>v1.2.0</b> — <i>Full release</i> — ?/?/? — Patch E</h2>

* Implement feedback from the ***gostack*** community
    * Bug fixes
    * Function restructuring
    * Documentational improvements
    * Conceptual improvements

---

 [> Return to glossary](../README.md)
