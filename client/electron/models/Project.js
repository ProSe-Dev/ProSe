const filesHelper = require("../helpers/files");
const fs = require("../helpers/fs");
const { resolve } = require("path");
const PROJECT_CONFIG_FOLDER_NAME = ".prose";
const SNAPSHOT_FILE_NAME = "snapshots.json";
const git = require("../helpers/git");
const Log = require("../helpers/log");
const constants = require("../../src/shared/constants");
const CONFIG_NAME = ".git/.prose";
const child = require("child_process").execFile;
var isWin = process.platform === "win32";
const settings = require("../settings");
const s = require("../../src/shared/settings");
const { appRoot } = require("../helpers/environment");
var hookFile = isWin ? "hook.exe" : "hook";
var postCommitFile = isWin ? "post-commit.exe" : "post-commit";
const inProduction = !process.env.DEVELOPMENT;

class Project {
  constructor(
    projectID,
    name,
    contact,
    abspath,
    creationDate,
    colorClass,
    isSynced,
    excludedFiles,
    snapshots,
    files
  ) {
    this.projectID = projectID;
    this.name = name;
    this.contact = contact;
    this.path = abspath;
    this.creationDate = creationDate;
    this.colorClass = colorClass;
    this.isSynced = isSynced || false;
    this.excludedFiles = excludedFiles || [];
    this.snapshots = snapshots || [];
    this.files = files || [];
  }

  async initialize() {
    Log.debugLog("Project.initialize", `initalizing`);
    let isExistingProject = await git.isGit(this.path);
    //Log.debugLog(isExistingProject);
    // TODO: add an option for this?
    this.isSynced = false;
    //this.isSynced = !isExistingProject;
    if (!isExistingProject) {
      await git.init(this.path);
    }
    await this.updateFiles();
    await this.writeConfig();
    // Hacky
    // TODO: figure out how this works when packaging
    let hookDes = resolve(this.path, ".git", "hooks", postCommitFile);
    let hookSrc;
    if (inProduction) {
      hookSrc = resolve(appRoot, "hook", hookFile);
    } else {
      hookSrc = resolve(__dirname, "..", "..", "hook", hookFile);
    }
    Log.debugLog(
      "Project.initialize",
      `Copying commit hook from ${hookSrc} to ${hookDes}`
    );
    fs.copyFileSync(hookSrc, hookDes);

    await this.commit();
  }

  getConfigPath() {
    return resolve(this.path, CONFIG_NAME);
  }

  async getTrackedFiles() {
    let fileLocations = [];
    for (let f of this.files) {
      if (f.status === constants.GIT_EXCLUDED) {
        continue;
      }
      fileLocations.push(f.path);
    }
    Log.debugLog("Got file locations: " + JSON.stringify(fileLocations));
    return fileLocations;
  }

  async writeConfig() {
    let projectInfo = {
      projectID: this.projectID,
      excludedFiles: this.excludedFiles,
      publicKey: settings.getVal(s.NAMESPACES.APP, s.KEYS.MASTER_KEYS)
        .publicKey,
      trackedFiles: await this.getTrackedFiles(),
      contact: this.contact,
      relayHost: settings.getVal(
        s.NAMESPACES.APP,
        s.KEYS.RELAY_HOST_ADDRESS,
        s.DEFAULTS.RELAY_HOST_ADDRESS
      )
    };
    Log.debugLog(JSON.stringify(projectInfo));
    await fs.writeFileAsync(this.getConfigPath(), JSON.stringify(projectInfo));
  }

  async updateFiles() {
    let fileStatuses = await git.projectStatus(this.path);
    Log.debugLog("updateFiles", "fileStatuses", fileStatuses);
    let files = fileStatuses;
    this.excludedFiles.forEach(exclude => {
      let existing = files.find(f => f.path === exclude);
      if (existing) {
        existing.status = constants.GIT_EXCLUDED;
      } else {
        files.push({
          path: exclude,
          status: constants.GIT_EXCLUDED
        });
      }
    });
    Log.debugLog("updateFiles", "files:", files);
    this.files = files;
  }

  async commit() {
    Log.debugLog("Project.commit", "COMMITING");
    let commitHash = await git.commit(this.path, this.excludedFiles);
    Log.debugLog("Project.commit", "Hash was: " + commitHash);
    if (!commitHash) {
      return;
    }
    this.snapshots.push({
      id: commitHash,
      creationDate: new Date()
    });
    await this.updateFiles();
    await this.writeConfig();
    Log.debugLog("Project.commit", "Completed commit");
    // Manually execute the git commit hook
    let executablePath = resolve(this.path, ".git", "hooks", postCommitFile);
    Log.debugLog("Project.commit", "Executing " + executablePath);
    child(executablePath, { cwd: this.path }, function(err, stdout, stderr) {
      if (err) {
        Log.debugLog("Project.commit", "ERROR: " + err);
        return;
      }
      Log.debugLog("Project.commit", "STDOUT: " + stdout.toString());
      Log.debugLog("Project.commit", "STDERR: " + stderr.toString());
    });
    await this.updateFiles();
  }

  async getFiles() {
    this.files = await filesHelper.getAllFilesInDirAsync(this.path);
    return this.files;
  }

  async getSnapshots() {
    if (!this.snapshots) {
      let snapshotsRaw = await fs.readFileAsync(
        resolve(this.path, PROJECT_CONFIG_FOLDER_NAME, SNAPSHOT_FILE_NAME)
      );
      this.snapshots = JSON.parse(snapshotsRaw);
    }
    return this.snapshots;
  }
}

module.exports = {
  Project
};
