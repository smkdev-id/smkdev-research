import 'dart:convert';

// TODO 1

class PersonalData {
  PersonalData({
    this.name,
    this.image,
    this.descSimple,
    this.descFull,
    this.socialMedia,
    this.skills,
    this.experiences,
    this.creations,
  });

  String name;
  String image;
  String descSimple;
  String descFull;
  List<SocialMedia> socialMedia;
  List<Skill> skills;
  List<Experience> experiences;
  List<Creation> creations;

  factory PersonalData.fromJson(String str) {
    return PersonalData.fromMap(json.decode(str));
  }

  String toJson() {
    return json.encode(toMap());
  }

  factory PersonalData.fromMap(Map<String, dynamic> json) {
    return PersonalData(
      name: json["name"],
      image: json["image"],
      descSimple: json["desc_simple"],
      descFull: json["desc_full"],
      socialMedia: List<SocialMedia>.from(json["social_media"].map((x) {
        return SocialMedia.fromMap(x);
      })),
      skills: List<Skill>.from(json["skills"].map((x) {
        return Skill.fromMap(x);
      })),
      experiences: List<Experience>.from(json["experiences"].map((x) {
        return Experience.fromMap(x);
      })),
      creations: List<Creation>.from(json["creations"].map((x) {
        return Creation.fromMap(x);
      })),
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "name": name,
      "image": image,
      "desc_simple": descSimple,
      "desc_full": descFull,
      "social_media": List<dynamic>.from(socialMedia.map((x) {
        return x.toMap();
      })),
      "skills": List<dynamic>.from(skills.map((x) {
        return x.toMap();
      })),
      "experiences": List<dynamic>.from(experiences.map((x) {
        return x.toMap();
      })),
      "creations": List<dynamic>.from(creations.map((x) {
        return x.toMap();
      })),
    };
  }
}

class Creation {
  Creation({
    this.title,
    this.url,
  });

  String title;
  String url;

  factory Creation.fromJson(String str) {
    return Creation.fromMap(json.decode(str));
  }

  String toJson() {
    return json.encode(toMap());
  }

  factory Creation.fromMap(Map<String, dynamic> json) {
    return Creation(
      title: json["title"],
      url: json["url"],
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "title": title,
      "url": url,
    };
  }
}

class Experience {
  Experience({
    this.title,
    this.period,
  });

  String title;
  String period;

  factory Experience.fromJson(String str) {
    return Experience.fromMap(json.decode(str));
  }

  String toJson() {
    return json.encode(toMap());
  }

  factory Experience.fromMap(Map<String, dynamic> json) {
    return Experience(
      title: json["title"],
      period: json["period"],
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "title": title,
      "period": period,
    };
  }
}

class Skill {
  Skill({
    this.title,
    this.rating,
  });

  String title;
  int rating;

  factory Skill.fromJson(String str) {
    return Skill.fromMap(json.decode(str));
  }

  String toJson() {
    return json.encode(toMap());
  }

  factory Skill.fromMap(Map<String, dynamic> json) {
    return Skill(
      title: json["title"],
      rating: json["rating"],
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "title": title,
      "rating": rating,
    };
  }
}

class SocialMedia {
  SocialMedia({
    this.title,
    this.image,
    this.url,
  });

  String title;
  String image;
  String url;

  factory SocialMedia.fromJson(String str) {
    return SocialMedia.fromMap(json.decode(str));
  }

  String toJson() {
    return json.encode(toMap());
  }

  factory SocialMedia.fromMap(Map<String, dynamic> json) {
    return SocialMedia(
      title: json["title"],
      image: json["image"],
      url: json["url"],
    );
  }

  Map<String, dynamic> toMap() {
    return {
      "title": title,
      "image": image,
      "url": url,
    };
  }
}
