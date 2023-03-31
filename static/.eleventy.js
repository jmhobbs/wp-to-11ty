// Generated by wp-to-11ty
const { DateTime } = require("luxon");
const slugify = require('@sindresorhus/slugify');
const util = require('util');

const categories = require('./_data/categories.json');
const post_tags = require('./_data/post_tags.json');

module.exports = function(eleventy) {
  eleventy.addPassthroughCopy("wp-content");

  eleventy.addFilter("strftime", function(dateObj, format) {
    return DateTime.fromJSDate(dateObj, {zone: 'utc'}).toFormat(format);
  });

  eleventy.addFilter('console', function(value) {
    return util.inspect(value);
  });

  eleventy.addFilter('wp_tag_slug', function (value) {
    return post_tags[value] || slugify(value);
  });

  eleventy.addFilter('wp_category_slug', function (value) {
    return categories[value]?.nice_name || slugify(value);
  });

  eleventy.addCollection("page", function (collections) {
    return collections.getAll().filter(function (item) {
      return "page" == item.data.type
    });
  });

  eleventy.addCollection("post", function (collections) {
    return collections.getAll().filter(function (item) {
      return "post" == item.data.type
    });
  });

  eleventy.addCollection("category", function (collections) {
    const categorized = {};
    collections.getAll().forEach(item => {
      if(item.data.category) {
        if(Array.isArray(item.data.category)) {
          item.data.category.forEach(category => {
            categorized[category] = categorized[category] || [];
            categorized[category].push(item);
          });
        } else {
          categorized[item.data.category] = categorized[item.data.category] || [];
          categorized[item.data.category].push(item);
        }
      }
    });
    return categorized;
  });

  eleventy.addCollection("categories", function (collections) {
    return collections.getAll()
      .map(function (item) { return item.data.category })
      .filter(function (category) { return !!(category); })
      .map(function (item) { return item[0] })
      .filter(function (category, index, arr) { return arr.indexOf(category) == index; });
  });

  return {};
};
