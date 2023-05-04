// Copyright (c) 2020 Mr.Coxall All rights reserved
//
// Created by: Alyssia fung
// Created on: May 2023
// This file contains the JS functions for index.html

"use strict"

/**
 * Check servie worker.
 */

function calculate() {
  let inputedAge = document.getElementById("age").value
  let inputedDay = document.getElementById("day").value

  if (
    inputedDay == "Tuesday" ||
    inputedDay == "Thursday" ||
    (inputedAge > 12 && inputedAge < 21)
  ) {
    document.getElementById("answer").innerHTML =
      "You're able to get the discount!"
  } else {
    document.getElementById("answer").innerHTML =
      "Sorry, you must pay for regular pricing."
  }
}
