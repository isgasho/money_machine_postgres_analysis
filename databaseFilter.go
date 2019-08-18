package main

import (
	"strings"
)

func filterDay() []string {
	arrayFiltered := []string{}
	return arrayFiltered
}

//entries within Hour within Day set
// function filterEntriesWithinHour(){

// }

//There's going to be entries within hour, minute, within a particular range set
func filterDowEntriesWithinTimeset(dataset []Dow, conditionRange1 string, conditionRange2 string) []Dow {
	var matchInRangeList []Dow

	splitYearCondition1 := strings.Split(conditionRange1, "T")[0]
	year1 := strings.Split(splitYearCondition1, "-")[0]

	splitMonthCondition1 := strings.Split(conditionRange1, "T")[0]
	month1 := strings.Split(splitMonthCondition1, "-")[1]

	splitDayCondition1 := strings.Split(conditionRange1, "T")[0]
	day1 := strings.Split(splitDayCondition1, "-")[2]

	splitHourCondition1 := strings.Split(conditionRange1, "T")[1]
	hour1 := strings.Split(splitHourCondition1, ":")[0]

	splitMinuteCondition1 := strings.Split(conditionRange1, "T")[1]
	minute1 := strings.Split(splitMinuteCondition1, ":")[1]

	splitSecondCondition1 := strings.Split(conditionRange1, "T")[1]
	second1 := strings.Split(splitSecondCondition1, ":")[2]

	//
	splitYearCondition2 := strings.Split(conditionRange2, "T")[0]
	year2 := strings.Split(splitYearCondition2, "-")[0]

	splitMonthCondition2 := strings.Split(conditionRange2, "T")[0]
	month2 := strings.Split(splitMonthCondition2, "-")[1]

	splitDayCondition2 := strings.Split(conditionRange2, "T")[0]
	day2 := strings.Split(splitDayCondition2, "-")[2]

	splitHourCondition2 := strings.Split(conditionRange2, "T")[1]
	hour2 := strings.Split(splitHourCondition2, ":")[0]

	splitMinuteCondition2 := strings.Split(conditionRange2, "T")[1]
	minute2 := strings.Split(splitMinuteCondition2, ":")[1]

	splitSecondCondition2 := strings.Split(conditionRange2, "T")[1]
	second2 := strings.Split(splitSecondCondition2, ":")[2]

	for i, v := range dataset {
		i++
		splitYear := strings.Split(v.CreatedAt, "T")[0]
		year := strings.Split(splitYear, "-")[0]

		splitMonth := strings.Split(v.CreatedAt, "T")[0]
		month := strings.Split(splitMonth, "-")[1]

		splitDay := strings.Split(v.CreatedAt, "T")[0]
		day := strings.Split(splitDay, "-")[2]

		splitHour := strings.Split(v.CreatedAt, "T")[1]
		hour := strings.Split(splitHour, ":")[0]

		splitMinute := strings.Split(v.CreatedAt, "T")[1]
		minute := strings.Split(splitMinute, ":")[1]

		splitSecond := strings.Split(v.CreatedAt, "T")[1]
		splitSecond1 := strings.Split(splitSecond, ":")[2]
		second := strings.Split(splitSecond1, ".")[0]

		//batch 1 case 1
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour == hour2 {
							if minute == minute2 {
								if second == second2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 1 case 2
		if year > year1 {
			if year < year2 {
				matchInRangeList = append(matchInRangeList, v)
				continue
			}
		}

		//batch 1 case 3
		if year > year1 {
			if year == year2 {
				if month < month2 {
					matchInRangeList = append(matchInRangeList, v)
					continue
				}
			}
		}

		//batch 1 case 4
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day < day2 {
						matchInRangeList = append(matchInRangeList, v)
						continue
					}
				}
			}
		}

		//batch 1 case 5
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour < hour2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 1 case 6
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour == hour2 {
							if minute < minute2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 1 case 7
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour == hour2 {
							if minute == minute2 {
								if second < second2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 2 case 1
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour == hour2 {
								if minute == minute2 {
									if second == second2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 2 case 2
		if year == year1 {
			if month > month1 {
				if year < year2 {
					matchInRangeList = append(matchInRangeList, v)
					continue
				}
			}
		}

		//batch 2 case 3
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month < month2 {
						matchInRangeList = append(matchInRangeList, v)
						continue
					}
				}
			}
		}

		//batch 2 case 4
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day < day2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 2 case 5
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour < hour2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 2 case 6
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour == hour2 {
								if minute < minute2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 2 case 7
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour == hour2 {
								if minute == minute2 {
									if second < second2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 1
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute == minute2 {
										if second == second2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 2
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year < year2 {
						matchInRangeList = append(matchInRangeList, v)
						continue
					}
				}
			}
		}

		//batch 3 case 3
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month < month2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 3 case 4
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day < day2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 3 case 5
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour < hour2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 6
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute < minute2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 7
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute == minute2 {
										if second < second2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		// fmt.Println(hour)
		// fmt.Println(hour2)
		// fmt.Println(minute)
		// fmt.Println(minute2)
		// fmt.Println(second)
		// fmt.Println(second2)

		//batch 3 case 8
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute == minute2 {
										if second == second2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour == hour2 {
										if minute == minute2 {
											if second == second2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 2
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year < year2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 4 case 3
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month < month2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 4 case 4
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day < day2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 5
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour < hour2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 6
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour == hour2 {
										if minute < minute2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 7
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour == hour2 {
										if minute == minute2 {
											if second < second2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour == hour2 {
											if minute == minute2 {
												if second == second2 {
													matchInRangeList = append(matchInRangeList, v)
													continue
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 2
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year < year2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 5 case 3
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month < month2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 4
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day < day2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 5
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour < hour2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 6
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour == hour2 {
											if minute < minute2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 7
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour == hour2 {
											if minute == minute2 {
												if second < second2 {
													matchInRangeList = append(matchInRangeList, v)
													continue
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour == hour2 {
												if minute == minute2 {
													if second == second2 {
														matchInRangeList = append(matchInRangeList, v)
														continue
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 2
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year < year2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 3
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month < month2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 4
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day < day2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 5
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour < hour2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 6
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour == hour2 {
												if minute < minute2 {
													matchInRangeList = append(matchInRangeList, v)
													continue
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 7
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour == hour2 {
												if minute == minute2 {
													if second < second2 {
														// fmt.Println("batch 6 case 7")
														matchInRangeList = append(matchInRangeList, v)
														continue
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 7 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second == second1 {
								matchInRangeList = append(matchInRangeList, v)
							}
						}
					}
				}
			}
		}
	}
	return matchInRangeList
}

func filterStockEntriesWithinTimeset(dataset []Stock, conditionRange1 string, conditionRange2 string) []Stock {
	var matchInRangeList []Stock

	splitYearCondition1 := strings.Split(conditionRange1, "T")[0]
	year1 := strings.Split(splitYearCondition1, "-")[0]

	splitMonthCondition1 := strings.Split(conditionRange1, "T")[0]
	month1 := strings.Split(splitMonthCondition1, "-")[1]

	splitDayCondition1 := strings.Split(conditionRange1, "T")[0]
	day1 := strings.Split(splitDayCondition1, "-")[2]

	splitHourCondition1 := strings.Split(conditionRange1, "T")[1]
	hour1 := strings.Split(splitHourCondition1, ":")[0]

	splitMinuteCondition1 := strings.Split(conditionRange1, "T")[1]
	minute1 := strings.Split(splitMinuteCondition1, ":")[1]

	splitSecondCondition1 := strings.Split(conditionRange1, "T")[1]
	second1 := strings.Split(splitSecondCondition1, ":")[2]

	//
	splitYearCondition2 := strings.Split(conditionRange2, "T")[0]
	year2 := strings.Split(splitYearCondition2, "-")[0]

	splitMonthCondition2 := strings.Split(conditionRange2, "T")[0]
	month2 := strings.Split(splitMonthCondition2, "-")[1]

	splitDayCondition2 := strings.Split(conditionRange2, "T")[0]
	day2 := strings.Split(splitDayCondition2, "-")[2]

	splitHourCondition2 := strings.Split(conditionRange2, "T")[1]
	hour2 := strings.Split(splitHourCondition2, ":")[0]

	splitMinuteCondition2 := strings.Split(conditionRange2, "T")[1]
	minute2 := strings.Split(splitMinuteCondition2, ":")[1]

	splitSecondCondition2 := strings.Split(conditionRange2, "T")[1]
	second2 := strings.Split(splitSecondCondition2, ":")[2]

	for i, v := range dataset {
		i++
		splitYear := strings.Split(v.CreatedAt, "T")[0]
		year := strings.Split(splitYear, "-")[0]

		splitMonth := strings.Split(v.CreatedAt, "T")[0]
		month := strings.Split(splitMonth, "-")[1]

		splitDay := strings.Split(v.CreatedAt, "T")[0]
		day := strings.Split(splitDay, "-")[2]

		splitHour := strings.Split(v.CreatedAt, "T")[1]
		hour := strings.Split(splitHour, ":")[0]

		splitMinute := strings.Split(v.CreatedAt, "T")[1]
		minute := strings.Split(splitMinute, ":")[1]

		splitSecond := strings.Split(v.CreatedAt, "T")[1]
		splitSecond1 := strings.Split(splitSecond, ":")[2]
		second := strings.Split(splitSecond1, ".")[0]

		//batch 1 case 1
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour == hour2 {
							if minute == minute2 {
								if second == second2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 1 case 2
		if year > year1 {
			if year < year2 {
				matchInRangeList = append(matchInRangeList, v)
				continue
			}
		}

		//batch 1 case 3
		if year > year1 {
			if year == year2 {
				if month < month2 {
					matchInRangeList = append(matchInRangeList, v)
					continue
				}
			}
		}

		//batch 1 case 4
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day < day2 {
						matchInRangeList = append(matchInRangeList, v)
						continue
					}
				}
			}
		}

		//batch 1 case 5
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour < hour2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 1 case 6
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour == hour2 {
							if minute < minute2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 1 case 7
		if year > year1 {
			if year == year2 {
				if month == month2 {
					if day == day2 {
						if hour == hour2 {
							if minute == minute2 {
								if second < second2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 2 case 1
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour == hour2 {
								if minute == minute2 {
									if second == second2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 2 case 2
		if year == year1 {
			if month > month1 {
				if year < year2 {
					matchInRangeList = append(matchInRangeList, v)
					continue
				}
			}
		}

		//batch 2 case 3
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month < month2 {
						matchInRangeList = append(matchInRangeList, v)
						continue
					}
				}
			}
		}

		//batch 2 case 4
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day < day2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 2 case 5
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour < hour2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 2 case 6
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour == hour2 {
								if minute < minute2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 2 case 7
		if year == year1 {
			if month > month1 {
				if year == year2 {
					if month == month2 {
						if day == day2 {
							if hour == hour2 {
								if minute == minute2 {
									if second < second2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 1
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute == minute2 {
										if second == second2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 2
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year < year2 {
						matchInRangeList = append(matchInRangeList, v)
						continue
					}
				}
			}
		}

		//batch 3 case 3
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month < month2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 3 case 4
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day < day2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 3 case 5
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour < hour2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 6
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute < minute2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 3 case 7
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute == minute2 {
										if second < second2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		// fmt.Println(hour)
		// fmt.Println(hour2)
		// fmt.Println(minute)
		// fmt.Println(minute2)
		// fmt.Println(second)
		// fmt.Println(second1)
		// fmt.Println(second2)

		//batch 3 case 8
		if year == year1 {
			if month == month1 {
				if day > day1 {
					if year == year2 {
						if month == month2 {
							if day == day2 {
								if hour == hour2 {
									if minute == minute2 {
										if second == second2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour == hour2 {
										if minute == minute2 {
											if second == second2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 2
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year < year2 {
							matchInRangeList = append(matchInRangeList, v)
							continue
						}
					}
				}
			}
		}

		//batch 4 case 3
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month < month2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 4 case 4
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day < day2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 5
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour < hour2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 6
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour == hour2 {
										if minute < minute2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 4 case 7
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour > hour1 {
						if year == year2 {
							if month == month2 {
								if day == day2 {
									if hour == hour2 {
										if minute == minute2 {
											if second < second2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour == hour2 {
											if minute == minute2 {
												if second == second2 {
													matchInRangeList = append(matchInRangeList, v)
													continue
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 2
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year < year2 {
								matchInRangeList = append(matchInRangeList, v)
								continue
							}
						}
					}
				}
			}
		}

		//batch 5 case 3
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month < month2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 4
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day < day2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 5
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour < hour2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 6
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour == hour2 {
											if minute < minute2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 5 case 7
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute > minute1 {
							if year == year2 {
								if month == month2 {
									if day == day2 {
										if hour == hour2 {
											if minute == minute2 {
												if second < second2 {
													matchInRangeList = append(matchInRangeList, v)
													continue
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour == hour2 {
												if minute == minute2 {
													if second == second2 {
														matchInRangeList = append(matchInRangeList, v)
														continue
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 2
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year < year2 {
									matchInRangeList = append(matchInRangeList, v)
									continue
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 3
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month < month2 {
										matchInRangeList = append(matchInRangeList, v)
										continue
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 4
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day < day2 {
											matchInRangeList = append(matchInRangeList, v)
											continue
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 5
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour < hour2 {
												matchInRangeList = append(matchInRangeList, v)
												continue
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 6
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour == hour2 {
												if minute < minute2 {
													matchInRangeList = append(matchInRangeList, v)
													continue
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 6 case 7
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second > second1 {
								if year == year2 {
									if month == month2 {
										if day == day2 {
											if hour == hour2 {
												if minute == minute2 {
													if second < second2 {
														matchInRangeList = append(matchInRangeList, v)
														continue
													}
												}
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}

		//batch 7 case 1
		if year == year1 {
			if month == month1 {
				if day == day1 {
					if hour == hour1 {
						if minute == minute1 {
							if second == second1 {
								matchInRangeList = append(matchInRangeList, v)
							}
						}
					}
				}
			}
		}
	}
	return matchInRangeList
}
