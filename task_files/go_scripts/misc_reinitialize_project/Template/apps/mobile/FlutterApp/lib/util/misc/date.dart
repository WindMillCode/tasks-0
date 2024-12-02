

extension DurationToJsonExtension on Duration {
  Map<String, dynamic> toJson() {
    return {
      'days': this.inDays,
      'hours': this.inHours % 24,
      'minutes': this.inMinutes % 60,
      'seconds': this.inSeconds % 60,
      'milliseconds': this.inMilliseconds % 1000,
      'microseconds': this.inMicroseconds % 1000,
    };
  }
}

extension DurationExtension on Duration {
  Map<String, int> normalize([int? value, String? unit]) {
    int totalMilliseconds = 0;

    if (value != null && unit != null) {
      switch (unit) {
        case 'years':
          totalMilliseconds = value * Duration.millisecondsPerDay * 365;
          break;
        case 'months':
          totalMilliseconds = value * Duration.millisecondsPerDay * 30;
          break;
        case 'days':
          totalMilliseconds = value * Duration.millisecondsPerDay;
          break;
        case 'hours':
          totalMilliseconds = value * Duration.millisecondsPerHour;
          break;
        case 'minutes':
          totalMilliseconds = value * Duration.millisecondsPerMinute;
          break;
        case 'seconds':
          totalMilliseconds = value * Duration.millisecondsPerSecond;
          break;
        case 'milliseconds':
          totalMilliseconds = value;
          break;
        default:
          throw ArgumentError('Invalid unit');
      }
    } else {
      totalMilliseconds = this.inMilliseconds;
    }

    int years = totalMilliseconds ~/ (Duration.millisecondsPerDay * 365);
    totalMilliseconds %= (Duration.millisecondsPerDay * 365);

    int months = totalMilliseconds ~/ (Duration.millisecondsPerDay * 30);
    totalMilliseconds %= (Duration.millisecondsPerDay * 30);

    int days = totalMilliseconds ~/ Duration.millisecondsPerDay;
    totalMilliseconds %= Duration.millisecondsPerDay;

    int hours = totalMilliseconds ~/ Duration.millisecondsPerHour;
    totalMilliseconds %= Duration.millisecondsPerHour;

    int minutes = totalMilliseconds ~/ Duration.millisecondsPerMinute;
    totalMilliseconds %= Duration.millisecondsPerMinute;

    int seconds = totalMilliseconds ~/ Duration.millisecondsPerSecond;
    totalMilliseconds %= Duration.millisecondsPerSecond;

    int milliseconds = totalMilliseconds;

    return {
      'years': years,
      'months': months,
      'days': days,
      'hours': hours,
      'minutes': minutes,
      'seconds': seconds,
      'milliseconds': milliseconds,
    };
  }
}
