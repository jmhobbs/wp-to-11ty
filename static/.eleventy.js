// Generated by wp-to-11ty
const { DateTime } = require("luxon");
const slugify = require('@sindresorhus/slugify');
const util = require('util');

const categories = require('./_data/categories.json');
const post_tags = require('./_data/post_tags.json');

module.exports = function(eleventy) {
  eleventy.addPassthroughCopy("wp-content");

  // Format a date/time with Luxon
  // {{ page.date | strftime("LLL d, y") }}
  eleventy.addFilter("strftime", function(dateObj, format) {
    return DateTime.fromJSDate(dateObj, {zone: 'utc'}).toFormat(format);
  });

  // Debugging tool, dumps the object being inspected
  // {{ page | console }}
  eleventy.addFilter('console', function(value) {
    return util.inspect(value);
  });

  // Get the slug/nice name of a tag, falling back to 11ty default
  // {{ tag | wp_tag_slug }}
  eleventy.addFilter('wp_tag_slug', function (value) {
    return post_tags[value] || slugify(value);
  });

  // Get the slug/nice name of a category, falling back to 11ty default
  // {{ category | wp_category_slug }}
  eleventy.addFilter('wp_category_slug', function (value) {
    return categories[value]?.nice_name || slugify(value);
  });

  // Dereference a variable from data based on input.
  // Mostly just a weird workaround for listing.njk
  eleventy.addFilter('ref', function (name) {
    return this.getVariables()[name];
  });

  // Get all pages by type
  eleventy.addCollection("page", function (collections) {
    return collections.getAll().filter(function (item) {
      return "page" == item.data.type
    });
  });

  // Get all posts by type
  eleventy.addCollection("post", function (collections) {
    return collections.getAll().filter(function (item) {
      return "page" != item.data.type
    });
  });

  // Get all entries, categorized.
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

  // Get a reference to all known categories.
  eleventy.addCollection("categories", function (collections) {
    return collections.getAll()
      .map(function (item) { return item.data.category })
      .filter(function (category) { return !!(category); })
      .map(function (item) { return item[0] })
      .filter(function (category, index, arr) { return arr.indexOf(category) == index; });
  });

  // Based on https://github.com/pdehaan/11ty-yearly-archives
  eleventy.addCollection("postsByYear", collection => {
    const data = {};

    collection.getAllSorted()
      .reverse()
      .filter(function (item) {
        return "page" != item.data.type
      })
      .forEach(post => {
        const year = post.date.getFullYear();
        const yearPosts = data[year] || [];
        yearPosts.push(post);
        data[year] = yearPosts;
      });

    return data;
  });

  return {};
};
