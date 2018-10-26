package cmd

import (
	"github.com/MegaShow/goagenda/controller"
	"github.com/spf13/cobra"
)

var meetingRootCmd = &cobra.Command{
	Use:     "meeting",
	Aliases: []string{"m", "meet"},
	Short:   "Meeting management",
	Long:    "Meeting management",
}

var meetingCreateCmd = &cobra.Command{
	Use:     "create",
	Aliases: []string{"c"},
	Short:   "Create a meeting",
	Long:    "Create a meeting",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetMeetingCtrl().MeetingCreate),
}

var meetingSetCmd = &cobra.Command{
	Use:     "set",
	Aliases: []string{"s"},
	Short:   "Set information of a meeting",
	Long:    "Set information of a meeting",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetMeetingCtrl().MeetingSet),
}

var meetingQuitCmd = &cobra.Command{
	Use:     "quit",
	Aliases: []string{"q"},
	Short:   "Quit a meeting",
	Long:    "Quit a meeting",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetMeetingCtrl().MeetingQuit),
}

var meetingDeleteCmd = &cobra.Command{
	Use:     "delete",
	Aliases: []string{"d"},
	Short:   "Delete your meeting",
	Long:    "Delete your meeting",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetMeetingCtrl().MeetingDelete),
}

var meetingAddCmd = &cobra.Command{
	Use:     "add",
	Aliases: []string{"a"},
	Short:   "Add your meeting participator",
	Long:    "Add your meeting participator",
	Args:    cobra.MinimumNArgs(1),
	Run:     wrapper(controller.GetMeetingCtrl().MeetingAdd),
}

var meetingListCmd = &cobra.Command{
	Use:     "list",
	Aliases: []string{"l"},
	Short:   "List your meeting",
	Long:    "List your meeting",
	Args:    cobra.NoArgs,
	Run:     wrapper(controller.GetMeetingCtrl().MeetingList),
}

var meetingRemoveCmd = &cobra.Command{
	Use:     "remove",
	Aliases: []string{"r"},
	Short:   "Remove your meeting participator",
	Long:    "Remove your meeting participator",
	Args:    cobra.MinimumNArgs(1),
	Run:     wrapper(controller.GetMeetingCtrl().MeetingRemove),
}

func init() {
	rootCmd.AddCommand(meetingRootCmd)

	meetingRootCmd.AddCommand(meetingCreateCmd)
	meetingRootCmd.AddCommand(meetingSetCmd)
	meetingRootCmd.AddCommand(meetingQuitCmd)
	meetingRootCmd.AddCommand(meetingDeleteCmd)
	meetingRootCmd.AddCommand(meetingAddCmd)
	meetingRootCmd.AddCommand(meetingListCmd)
	meetingRootCmd.AddCommand(meetingRemoveCmd)

	meetingCreateCmd.Flags().StringP("title", "t", "", "title of new meeting")
	meetingCreateCmd.Flags().StringP("startTime", "s", "", "start time of new meeting")
	meetingCreateCmd.Flags().StringP("endTime", "e", "", "end time of new meeting")
	meetingCreateCmd.Flags().StringSliceP("participator", "p", []string{}, "participators of new meeting")
	meetingCreateCmd.MarkFlagRequired("title")
	meetingCreateCmd.MarkFlagRequired("startTime")
	meetingCreateCmd.MarkFlagRequired("endTime")
	meetingCreateCmd.MarkFlagRequired("participator")

	meetingSetCmd.Flags().StringP("title", "t", "", "title of new meeting")
	meetingSetCmd.Flags().StringP("startTime", "s", "", "start time of new meeting")
	meetingSetCmd.Flags().StringP("endTime", "e", "", "end time of new meeting")
	meetingSetCmd.Flags().StringSliceP("participator", "p", []string{}, "participators of new meeting")
	meetingSetCmd.MarkFlagRequired("title")

	meetingQuitCmd.Flags().StringP("title", "t", "", "title of new meeting")
	meetingQuitCmd.MarkFlagRequired("title")

	meetingListCmd.Flags().StringP("title", "t", "", "the meeting title")
	meetingListCmd.Flags().StringP("startTime", "s", "", "the meeting start time")
	meetingListCmd.Flags().StringP("endTime", "e", "", "the meeting end time")
	meetingListCmd.Flags().StringP("userName", "u", "", "the user name of the host of the meeting")
	meetingListCmd.MarkFlagRequired("title")

	meetingDeleteCmd.Flags().StringP("title", "t", "", "new password")
	meetingDeleteCmd.Flags().BoolP("all", "a", false, "delete all meeting")

	meetingAddCmd.Flags().StringP("title", "t", "", "meeting title")
	meetingAddCmd.MarkFlagRequired("title")

	meetingRemoveCmd.Flags().StringP("title", "t", "", "meeting title")
	meetingRemoveCmd.MarkFlagRequired("title")
}
