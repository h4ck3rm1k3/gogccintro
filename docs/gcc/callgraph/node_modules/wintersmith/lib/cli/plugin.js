(function() {
  var NpmAdapter, async, chalk, clip, commonOptions, displayListing, extendOptions, fetchListing, fileExists, fs, https, loadEnv, logger, lpad, main, max, mkdirp, normalizePluginName, npm, options, path, readJSON, ref, ref1, usage, waterfall,
    indexOf = [].indexOf || function(item) { for (var i = 0, l = this.length; i < l; i++) { if (i in this && this[i] === item) return i; } return -1; };

  async = require('async');

  chalk = require('chalk');

  fs = require('fs');

  path = require('path');

  npm = require('npm');

  mkdirp = require('mkdirp');

  https = require('https');

  ref = require('./common'), NpmAdapter = ref.NpmAdapter, loadEnv = ref.loadEnv, commonOptions = ref.commonOptions, extendOptions = ref.extendOptions;

  ref1 = require('./../core/utils'), fileExists = ref1.fileExists, readJSON = ref1.readJSON;

  logger = require('./../core/logger').logger;

  usage = "\nusage: wintersmith plugin [options] <command>\n\ncommands:\n\n  " + (chalk.bold('list')) + " - list available plugins\n  " + (chalk.bold('install')) + " <plugin> - install plugin\n\noptions:\n\n  -C, --chdir [path]      change the working directory\n  -c, --config [path]     path to config\n";

  options = {};

  extendOptions(options, commonOptions);

  max = function(array, get) {
    var item, j, len, rv, v;
    if (get == null) {
      get = function(item) {
        return item;
      };
    }
    rv = null;
    for (j = 0, len = array.length; j < len; j++) {
      item = array[j];
      v = get(item);
      if (v > rv) {
        rv = v;
      }
    }
    return rv;
  };

  lpad = function(string, amount, char) {
    var i, j, p, ref2;
    if (char == null) {
      char = ' ';
    }
    p = '';
    for (i = j = 0, ref2 = amount - string.length; 0 <= ref2 ? j < ref2 : j > ref2; i = 0 <= ref2 ? ++j : --j) {
      p += char;
    }
    return p + string;
  };

  clip = function(string, maxlen) {
    if (string.length <= maxlen) {
      return string;
    }
    return string.slice(0, maxlen - 2).trim() + "..";
  };

  fetchListing = function(callback) {
    var request;
    return request = https.get('https://api.npms.io/v2/search?q=keywords:wintersmith-plugin&size=200', function(response) {
      var data, error;
      if (response.statusCode !== 200) {
        error = new Error("Unexpected response when searching registry, HTTP " + response.statusCode);
      }
      if (!/^application\/json/.test(response.headers['content-type'])) {
        error = new Error("Invalid content-type: " + response.headers['content-type']);
      }
      if (error != null) {
        response.resume();
        callback(error);
        return;
      }
      data = [];
      response.on('data', function(chunk) {
        return data.push(chunk);
      });
      return response.on('end', function() {
        var listing, parsed;
        try {
          parsed = JSON.parse(Buffer.concat(data));
        } catch (error1) {
          error = error1;
          callback(error);
          return;
        }
        listing = parsed.results.map(function(result) {
          return result["package"];
        });
        listing.sort(function(a, b) {
          if (a.name > b.name) {
            return 1;
          }
          if (a.name < b.name) {
            return -1;
          }
          return 0;
        });
        return callback(null, listing);
      });
    });
  };

  displayListing = function(list, callback) {
    var display, j, k, left, len, line, margin, maxw, pad, plugin, results;
    display = list.map(function(plugin) {
      var description, homepage, maintainers, name, ref2, ref3, ref4;
      name = normalizePluginName(plugin.name);
      description = plugin.description;
      maintainers = plugin.maintainers.map(function(v) {
        return v.username;
      }).join(' ');
      homepage = (ref2 = (ref3 = plugin.links) != null ? ref3.homepage : void 0) != null ? ref2 : (ref4 = plugin.links) != null ? ref4.npm : void 0;
      return {
        name: name,
        description: description,
        maintainers: maintainers,
        homepage: homepage
      };
    });
    pad = max(display, function(item) {
      return item.name.length;
    });
    maxw = process.stdout.getWindowSize()[0] - 2;
    margin = ((function() {
      results = [];
      for (var j = 0; 0 <= pad ? j < pad : j > pad; 0 <= pad ? j++ : j--){ results.push(j); }
      return results;
    }).apply(this).map(function() {
      return ' ';
    })).join('');
    for (k = 0, len = display.length; k < len; k++) {
      plugin = display[k];
      line = (lpad(plugin.name, pad)) + "  " + (clip(plugin.description, maxw - pad - 2));
      left = maxw - line.length;
      if (left > plugin.maintainers.length) {
        line += chalk.grey(lpad(plugin.maintainers, left));
      }
      logger.info(line.replace(/^\s*(\S+)  /, function(m) {
        return chalk.bold(m);
      }));
      if ((plugin.homepage != null) && plugin.homepage.length < maxw - pad - 2) {
        logger.info(margin + "  " + (chalk.gray(plugin.homepage)));
      }
      logger.info('');
    }
    return callback(null, list);
  };

  waterfall = function(flow, callback) {

    /* async.waterfall that allows for parallel tasks */
    var item, j, len, resolved;
    resolved = [];
    for (j = 0, len = flow.length; j < len; j++) {
      item = flow[j];
      switch (typeof item) {
        case 'function':
          resolved.push(item);
          break;
        case 'object':
        case 'array':
          resolved.push(async.apply(async.parallel, item));
          break;
        default:
          return callback(new Error("Invalid item '" + item + "' in flow"));
      }
    }
    return async.waterfall(resolved, callback);
  };

  normalizePluginName = function(name) {
    return name.replace(/^wintersmith\-/, '');
  };

  main = function(argv) {
    var action, cmd, installPlugin, loadCurrentEnv;
    action = argv._[3];
    if (action == null) {
      console.log(usage);
      process.exit(0);
    }
    loadCurrentEnv = function(callback) {
      return loadEnv(argv, callback);
    };
    installPlugin = function(res, callback) {
      var configFile, createPackageJson, env, install, j, len, list, name, p, packageFile, plugin, readConfig, saveConfig, updateConfig;
      env = res[0], list = res[1];
      name = argv._[4];
      plugin = null;
      for (j = 0, len = list.length; j < len; j++) {
        p = list[j];
        if (normalizePluginName(p.name) === normalizePluginName(name)) {
          plugin = p;
          break;
        }
      }
      if (!plugin) {
        return callback(new Error("Unknown plugin: " + name));
      }
      configFile = env.config.__filename;
      packageFile = env.resolvePath('package.json');
      createPackageJson = function(callback) {
        return fileExists(packageFile, function(exists) {
          if (exists) {
            return callback();
          } else {
            logger.warn("package.json missing, creating minimal package");
            return fs.writeFile(packageFile, '{\n  "dependencies": {},\n  "private": true\n}\n', callback);
          }
        });
      };
      readConfig = function(callback) {
        return readJSON(configFile, callback);
      };
      updateConfig = function(config, callback) {
        var ref2;
        if (config.plugins == null) {
          config.plugins = [];
        }
        if (ref2 = plugin.name, indexOf.call(config.plugins, ref2) < 0) {
          config.plugins.push(plugin.name);
        }
        return callback(null, config);
      };
      saveConfig = function(config, callback) {
        var json;
        logger.verbose("saving config file: " + configFile);
        json = JSON.stringify(config, null, 2);
        return fs.writeFile(configFile, json + '\n', callback);
      };
      install = function(callback) {
        logger.verbose("installing " + plugin.name);
        process.chdir(env.workDir);
        return async.series([
          createPackageJson, function(callback) {
            return npm.load({
              logstream: new NpmAdapter(logger),
              save: true
            }, callback);
          }, function(callback) {
            return npm.install(plugin.name, callback);
          }
        ], function(error) {
          return callback(error);
        });
      };
      return async.waterfall([install, readConfig, updateConfig, saveConfig], callback);
    };
    switch (action) {
      case 'list':
        cmd = [fetchListing, displayListing];
        break;
      case 'install':
        cmd = [[loadCurrentEnv, fetchListing], installPlugin];
        break;
      default:
        cmd = [
          function(callback) {
            return callback(new Error("Unknown plugin action: " + action));
          }
        ];
    }
    return waterfall(cmd, function(error) {
      if (error != null) {
        logger.error(error.message, error);
        return process.exit(1);
      } else {
        return process.exit(0);
      }
    });
  };

  module.exports = main;

  module.exports.usage = usage;

  module.exports.options = options;

}).call(this);
