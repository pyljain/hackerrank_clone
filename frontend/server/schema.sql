CREATE TABLE IF NOT EXISTS question (
  id serial primary key,
  title text,
  description text,
  category varchar(50),
  no_of_test_cases integer
);

CREATE TABLE IF NOT EXISTS submission (
  id serial primary key,
  userId text,
  qstnId integer REFERENCES question(id),
  solution text,
  language varchar(50),
  status text
);

CREATE TABLE IF NOT EXISTS submission_outcomes (
  id serial primary key,
  test_case text,
  expected_outcome text,
  actual_outcome text,
  submission_id integer REFERENCES submission(id)
);

INSERT INTO question (title, description, category) VALUES (
  'Hello World',
  'Complete the function that outputs hello world',
  'Easy'
);

INSERT INTO question (title, description, category) VALUES (
  'Simple Array Sum',
  'Given an array of integers, find the sum of its elements.\nFor example, if the array ar=[1, 2, 3], 1 + 2 + 3 = 6, so return 6.',
  'Easy'
);

INSERT INTO question (title, description, category) VALUES (
  'Birthday Cake Candles',
  'You are in charge of the cake for your nieces birthday and have decided the cake will have one candle for each year of her total age. When she blows out the candles, she’ll only be able to blow out the tallest ones. Your task is to find out how many candles she can successfully blow out. For example, if your niece is turning  years old, and the cake will have  candles of height , , , , she will be able to blow out  candles successfully, since the tallest candles are of height  and there are  such candles.',
  'Easy'
);

INSERT INTO question (title, description, category) VALUES (
  'Mini-Max Sum',
  'Given five positive integers, find the minimum and maximum values that can be calculated by summing exactly four of the five integers. Then print the respective minimum and maximum values as a single line of two space-separated long integers. For example, . Our minimum sum is  and our maximum sum is . We would print',
  'Medium'
);

INSERT INTO question (title, description, category) VALUES (
  'Time Conversion',
  'Given a time in 12-hour AM/PM format, convert it to military (24-hour) time. Note: Midnight is 12:00:00AM on a 12-hour clock, and 00:00:00 on a 24-hour clock. Noon is 12:00:00PM on a 12-hour clock, and 12:00:00 on a 24-hour clock.',
  'Hard'
);