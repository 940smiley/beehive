/*
 *    Copyright (C) 2014-2017 Christian Muehlhaeuser
 *
 *    This program is free software: you can redistribute it and/or modify
 *    it under the terms of the GNU Affero General Public License as published
 *    by the Free Software Foundation, either version 3 of the License, or
 *    (at your option) any later version.
 *
 *    This program is distributed in the hope that it will be useful,
 *    but WITHOUT ANY WARRANTY; without even the implied warranty of
 *    MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 *    GNU Affero General Public License for more details.
 *
 *    You should have received a copy of the GNU Affero General Public License
 *    along with this program.  If not, see <http://www.gnu.org/licenses/>.
 *
 *    Authors:
 *      Christian Muehlhaeuser <muesli@gmail.com>
 */

package huebee

import (
	"github.com/muesli/beehive/bees"
)

// HueBeeFactory is a factory for HueBees.
type HueBeeFactory struct {
	bees.BeeFactory
}

// New returns a new Bee instance configured with the supplied options.
func (factory *HueBeeFactory) New(name, description string, options bees.BeeOptions) bees.BeeInterface {
	bee := HueBee{
		Bee: bees.NewBee(name, factory.ID(), description, options),
	}
	bee.ReloadOptions(options)

	return &bee
}

// ID returns the ID of this Bee.
func (factory *HueBeeFactory) ID() string {
	return "huebee"
}

// Name returns the name of this Bee.
func (factory *HueBeeFactory) Name() string {
	return "Philips Hue"
}

// Description returns the description of this Bee.
func (factory *HueBeeFactory) Description() string {
	return "Controls Philips Hue lighting systems"
}

// Image returns the filename of an image for this Bee.
func (factory *HueBeeFactory) Image() string {
	return factory.ID() + ".png"
}

// LogoColor returns the preferred logo background color (used by the admin interface).
func (factory *HueBeeFactory) LogoColor() string {
	return "#212727"
}

// Options returns the options available to configure this Bee.
func (factory *HueBeeFactory) Options() []bees.BeeOptionDescriptor {
	opts := []bees.BeeOptionDescriptor{
		{
			Name:        "address",
			Description: "Address of the Hue bridge, eg: 192.168.0.1",
			Type:        "address",
			Mandatory:   true,
		},
		{
			Name:        "key",
			Description: "Key used for auth with the bridge",
			Type:        "string",
			Mandatory:   true,
		},
	}
	return opts
}

// Events describes the available events provided by this Bee.
func (factory *HueBeeFactory) Events() []bees.EventDescriptor {
	events := []bees.EventDescriptor{}
	return events
}

// Actions describes the available actions provided by this Bee.
func (factory *HueBeeFactory) Actions() []bees.ActionDescriptor {
	actions := []bees.ActionDescriptor{
		{
			Namespace:   factory.Name(),
			Name:        "switch",
			Description: "Switches on/off a Hue light",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "light",
					Description: "ID of the light you want to switch on or off",
					Type:        "int",
					Mandatory:   true,
				},
				{
					Name:        "state",
					Description: "New state of the light, true for turning it on",
					Type:        "bool",
					Mandatory:   true,
				},
			},
		},
		{
			Namespace:   factory.Name(),
			Name:        "setcolor",
			Description: "Changes the color of a Hue light",
			Options: []bees.PlaceholderDescriptor{
				{
					Name:        "light",
					Description: "ID of the light you want to switch on or off",
					Type:        "int",
					Mandatory:   true,
				},
				{
					Name:        "color",
					Description: "New color of the light you want to change",
					Type:        "string",
					Mandatory:   false,
				},
				{
					Name:        "brightness",
					Description: "New brightness of the light you want to change",
					Type:        "int",
					Mandatory:   false,
				},
				{
					Name:        "alert",
					Description: "0: no alert, 1: short alert, 2: long alert",
					Type:        "int",
					Mandatory:   false,
				},
			},
		},
	}
	return actions
}

func init() {
	f := HueBeeFactory{}
	bees.RegisterFactory(&f)
}
